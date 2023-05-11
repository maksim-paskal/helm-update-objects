package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	rspb "helm.sh/helm/v3/pkg/release"
)

var b64 = base64.StdEncoding

var magicGzip = []byte{0x1f, 0x8b, 0x08}

// encodeRelease encodes a release returning a base64 encoded
// gzipped string representation, or error.
func EncodeRelease(rls *rspb.Release) (string, error) {
	b, err := json.Marshal(rls)
	if err != nil {
		return "", errors.Wrap(err, "error marshaling release")
	}

	var buf bytes.Buffer

	w, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return "", errors.Wrap(err, "error creating gzip writer")
	}

	if _, err = w.Write(b); err != nil {
		return "", errors.Wrap(err, "error writing to gzip writer")
	}

	w.Close()

	return b64.EncodeToString(buf.Bytes()), nil
}

// decodeRelease decodes the bytes of data into a release
// type. Data must contain a base64 encoded gzipped string of a
// valid release, otherwise an error is returned.
func DecodeRelease(data string) (*rspb.Release, error) {
	// base64 decode string
	b, err := b64.DecodeString(data)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding release")
	}

	// For backwards compatibility with releases that were stored before
	// compression was introduced we skip decompression if the
	// gzip magic header is not found
	if len(b) > 3 && bytes.Equal(b[0:3], magicGzip) {
		r, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			return nil, errors.Wrap(err, "error creating gzip reader")
		}

		defer r.Close()

		b2, err := io.ReadAll(r)
		if err != nil {
			return nil, errors.Wrap(err, "error reading gzip data")
		}

		b = b2
	}

	var rls rspb.Release
	// unmarshal release object bytes
	if err := json.Unmarshal(b, &rls); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling release")
	}

	return &rls, nil
}
