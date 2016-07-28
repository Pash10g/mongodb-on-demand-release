package adapter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func om_cluster_docs_replica_set_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x57,
		0x4d, 0x73, 0xa3, 0x38, 0x10, 0xbd, 0xfb, 0x57, 0xb8, 0xd8, 0xab, 0x3f,
		0x88, 0xe3, 0x64, 0x93, 0x9c, 0xc6, 0x93, 0x1c, 0xe6, 0xb0, 0xd9, 0x4c,
		0x25, 0x33, 0x93, 0xc3, 0x4c, 0xca, 0x25, 0x90, 0x30, 0xda, 0x11, 0x12,
		0x25, 0x89, 0xc4, 0x0e, 0xc5, 0x7f, 0x5f, 0x49, 0x60, 0x8c, 0x40, 0x38,
		0x9e, 0xaa, 0xf1, 0x66, 0x39, 0x99, 0x56, 0xab, 0xbb, 0xf5, 0xf4, 0x5e,
		0xd3, 0xce, 0x07, 0x43, 0xf5, 0x78, 0x2c, 0x95, 0x98, 0x51, 0xe1, 0x5d,
		0x0d, 0x4b, 0x83, 0x31, 0x42, 0xf6, 0x42, 0x09, 0x03, 0xf0, 0x23, 0x10,
		0x48, 0xad, 0x78, 0xd3, 0x67, 0xc0, 0xa7, 0x04, 0x07, 0xd3, 0x84, 0xd1,
		0x15, 0x83, 0xc1, 0x38, 0x49, 0xc4, 0x18, 0x64, 0x92, 0x25, 0x40, 0xef,
		0xf6, 0x46, 0xee, 0xad, 0x8f, 0x98, 0xaa, 0x57, 0x1d, 0xdb, 0xbb, 0xbe,
		0xfa, 0xf1, 0xa3, 0x67, 0xb3, 0xd9, 0x5b, 0x94, 0x21, 0x3c, 0xe3, 0x73,
		0x13, 0x7c, 0x43, 0x5c, 0x54, 0x65, 0x7d, 0x6f, 0xd4, 0x15, 0x64, 0x98,
		0x40, 0x63, 0xac, 0x6d, 0xdb, 0x27, 0xef, 0x58, 0xca, 0x1d, 0x58, 0x6a,
		0xff, 0xf3, 0xf9, 0xc8, 0xbd, 0x1e, 0x11, 0xf0, 0xcc, 0xb8, 0xae, 0xd0,
		0xeb, 0xf1, 0x58, 0x61, 0x59, 0x55, 0xa3, 0xbd, 0xe6, 0xb3, 0xf9, 0x65,
		0x78, 0x02, 0x67, 0xc1, 0xd9, 0xe5, 0xe5, 0x25, 0x0a, 0x82, 0xe8, 0x24,
		0x82, 0xd1, 0x69, 0x10, 0xfa, 0xc8, 0x47, 0xa7, 0xc1, 0x69, 0x14, 0x9d,
		0x85, 0x3e, 0x00, 0xd1, 0xac, 0x2f, 0x5a, 0x02, 0xd6, 0x77, 0xa2, 0x11,
		0xaf, 0xd7, 0x0f, 0xd3, 0xc3, 0xfc, 0x18, 0xcc, 0x08, 0x32, 0x90, 0x3c,
		0xf5, 0xb8, 0xa4, 0x04, 0xc8, 0x88, 0xf1, 0x44, 0x87, 0x61, 0x62, 0xdd,
		0x17, 0x29, 0xe3, 0x44, 0x7b, 0xc4, 0x52, 0xa6, 0xe2, 0x6a, 0x3a, 0x8d,
		0x80, 0x90, 0x90, 0x4c, 0xaa, 0x3b, 0x9b, 0x30, 0xbe, 0x52, 0x04, 0xa0,
		0xd9, 0xba, 0xa6, 0x80, 0x79, 0x1b, 0xaf, 0x2f, 0xce, 0x97, 0xe7, 0xf3,
		0x31, 0x48, 0xc0, 0x2b, 0xa3, 0xe3, 0xd3, 0xc9, 0x6c, 0xf2, 0xe7, 0x44,
		0xae, 0x5e, 0xfb, 0x92, 0xbc, 0x60, 0x3a, 0xf3, 0xfd, 0x8b, 0x94, 0x64,
		0xba, 0xe4, 0x08, 0x10, 0x81, 0xfa, 0x3d, 0xbf, 0x5d, 0xdf, 0x23, 0x88,
		0x85, 0xbc, 0x21, 0x64, 0x1f, 0x04, 0x0d, 0xcf, 0xbb, 0x9a, 0xcc, 0xbd,
		0x68, 0x34, 0xbc, 0xbf, 0xf2, 0x43, 0xe3, 0x36, 0x2f, 0xa2, 0xe3, 0x5c,
		0x58, 0x96, 0x46, 0x5e, 0x8f, 0x82, 0xc4, 0xc8, 0xc7, 0xc0, 0x52, 0xf1,
		0xbc, 0x5a, 0xf7, 0x02, 0x10, 0xfe, 0xcc, 0xd2, 0x1e, 0x9e, 0xc7, 0x4c,
		0xc8, 0xed, 0xe6, 0x3c, 0xa7, 0x0c, 0x22, 0x31, 0xf9, 0xee, 0x3f, 0x15,
		0x45, 0x53, 0x69, 0x84, 0xad, 0x3e, 0x03, 0x19, 0xef, 0xf4, 0xc9, 0x56,
		0x3d, 0xfa, 0x9c, 0x96, 0xc9, 0xc6, 0x60, 0x85, 0xa8, 0x9c, 0x28, 0xbf,
		0x56, 0x98, 0x7b, 0x26, 0x81, 0x44, 0x56, 0x0b, 0x30, 0x4b, 0x02, 0xbf,
		0xa2, 0x2f, 0x31, 0x47, 0x22, 0x66, 0x04, 0xde, 0x7e, 0x54, 0x0e, 0x27,
		0xbe, 0xef, 0xdb, 0x70, 0x79, 0x12, 0x27, 0x3b, 0xa7, 0x4f, 0x5c, 0x1f,
		0x65, 0x36, 0x1f, 0xd8, 0xe0, 0xe8, 0x53, 0xd7, 0xfa, 0xc6, 0x92, 0x71,
		0x4c, 0x57, 0xff, 0xd1, 0xd1, 0x77, 0x09, 0xdf, 0xf7, 0xf8, 0xc6, 0x39,
		0xe5, 0x2c, 0x44, 0x42, 0x94, 0x7a, 0xcd, 0xf3, 0x3f, 0x10, 0x08, 0xe3,
		0xa1, 0x39, 0x63, 0xb1, 0x65, 0x51, 0x03, 0x0a, 0xc0, 0x57, 0x62, 0xb6,
		0x3c, 0xef, 0x16, 0x46, 0x91, 0xec, 0x18, 0xcb, 0xf8, 0x8c, 0xeb, 0x95,
		0xd9, 0x85, 0xaa, 0xd3, 0x5a, 0x2d, 0x5a, 0x55, 0x73, 0x94, 0x12, 0x1c,
		0x96, 0xfd, 0xd7, 0x19, 0x49, 0x3b, 0x3c, 0x20, 0xf9, 0x77, 0x75, 0x13,
		0x69, 0x18, 0x2d, 0xb5, 0xc9, 0xdb, 0x1b, 0x55, 0x28, 0xa0, 0x15, 0xc8,
		0xee, 0x88, 0x30, 0xb0, 0xae, 0xec, 0x39, 0x04, 0xe9, 0x54, 0x6f, 0x40,
		0xf5, 0xcd, 0x41, 0x20, 0xc1, 0x1b, 0x09, 0x36, 0x42, 0xa2, 0xe4, 0x2f,
		0x75, 0x83, 0xee, 0x14, 0x48, 0x48, 0x4c, 0xb7, 0xa7, 0xf2, 0x22, 0x4c,
		0x90, 0x43, 0xdc, 0x5e, 0xda, 0xad, 0x63, 0x23, 0x1a, 0x1c, 0x5a, 0xea,
		0x1b, 0xd9, 0x56, 0x65, 0xf8, 0x62, 0x17, 0x35, 0x70, 0x94, 0xd7, 0xe2,
		0xad, 0x8c, 0xb1, 0xe8, 0x50, 0xf6, 0x98, 0x3c, 0x73, 0x74, 0x1d, 0x57,
		0x15, 0x15, 0x01, 0xbf, 0x6c, 0x52, 0xe3, 0x52, 0x9e, 0xb1, 0xe9, 0xf0,
		0xbc, 0xeb, 0x74, 0x65, 0xd7, 0x6a, 0xac, 0x29, 0x55, 0xc5, 0x0f, 0x61,
		0x8c, 0x12, 0xb0, 0xeb, 0x87, 0x67, 0x25, 0xc1, 0x15, 0x97, 0x71, 0x34,
		0xfc, 0x40, 0xd4, 0x47, 0xa3, 0x50, 0x2f, 0x48, 0xf5, 0xf5, 0xa2, 0x18,
		0xe5, 0xf9, 0x14, 0x47, 0x15, 0xb3, 0xd5, 0x6f, 0xcd, 0x76, 0xf3, 0xb6,
		0x95, 0x43, 0xc5, 0x42, 0xc5, 0xb3, 0x76, 0x17, 0x58, 0x62, 0x68, 0xd1,
		0xae, 0x51, 0x45, 0x82, 0x92, 0x00, 0xf1, 0xf6, 0x10, 0xe0, 0x16, 0x93,
		0x59, 0xb1, 0x71, 0x2c, 0x23, 0xe7, 0xf9, 0x07, 0x35, 0x9b, 0xa0, 0x75,
		0xd1, 0x66, 0x98, 0xbe, 0x46, 0x0b, 0x3c, 0x3b, 0xc7, 0xee, 0x90, 0x23,
		0x25, 0x4f, 0x35, 0x5c, 0x20, 0x7e, 0x47, 0xc9, 0x46, 0xed, 0x90, 0x3c,
		0x43, 0x23, 0x85, 0x2f, 0x56, 0xbd, 0x46, 0x6a, 0x83, 0x6f, 0x15, 0x51,
		0x41, 0x52, 0x63, 0xd2, 0x84, 0xa3, 0x7c, 0x9e, 0xec, 0x56, 0xc1, 0x99,
		0xfd, 0x59, 0xf7, 0x44, 0x0c, 0x38, 0x54, 0x6d, 0xac, 0xb2, 0x0d, 0xea,
		0x2b, 0xb1, 0x47, 0x37, 0xdd, 0xfa, 0xbe, 0x0a, 0x64, 0x46, 0x9a, 0xfe,
		0x29, 0x4d, 0x5b, 0x3f, 0xbf, 0xc0, 0xf2, 0xa4, 0xc3, 0x14, 0x08, 0xf1,
		0xc2, 0x38, 0x1c, 0xda, 0x64, 0x81, 0x0a, 0x79, 0xb6, 0x49, 0x54, 0xd3,
		0x5c, 0xa8, 0x34, 0xb7, 0x28, 0x8c, 0x01, 0xc5, 0x22, 0xe9, 0xce, 0x5f,
		0xde, 0xc3, 0xf5, 0xfd, 0xe2, 0x76, 0xfc, 0xf0, 0x69, 0x31, 0x3e, 0xd9,
		0x01, 0xd6, 0xfc, 0x12, 0xfe, 0x44, 0x9b, 0x2a, 0x97, 0xfa, 0xd5, 0x4a,
		0xa3, 0x2c, 0x46, 0xa9, 0x96, 0x22, 0xff, 0x61, 0x81, 0xb0, 0xe4, 0x18,
		0x32, 0x1a, 0xe1, 0x4a, 0xa1, 0x4b, 0x96, 0x4c, 0x74, 0xc4, 0x66, 0xad,
		0x58, 0x80, 0x80, 0x20, 0xd8, 0x9d, 0x2a, 0xbc, 0x4c, 0xa1, 0x21, 0x6e,
		0x10, 0x41, 0xd2, 0x2c, 0x37, 0x27, 0x83, 0x72, 0xed, 0x11, 0xd0, 0x6a,
		0xc9, 0x3a, 0x95, 0xb3, 0x87, 0xe9, 0x2a, 0x01, 0x54, 0xd3, 0x99, 0xab,
		0xb1, 0xd4, 0x77, 0xe6, 0x1c, 0x28, 0xdc, 0x13, 0xea, 0x01, 0x71, 0xad,
		0xf8, 0xda, 0x31, 0x54, 0xe3, 0x93, 0x62, 0xde, 0x6d, 0xf9, 0x69, 0xeb,
		0x4e, 0x24, 0xfa, 0x29, 0x3a, 0x56, 0xc7, 0x44, 0x64, 0xce, 0xbf, 0x65,
		0x4a, 0xfb, 0x4b, 0xe9, 0x3a, 0x21, 0x56, 0x2e, 0x6e, 0xde, 0xec, 0xeb,
		0xdb, 0xff, 0x7b, 0x24, 0x17, 0xc6, 0xdf, 0x8d, 0xa3, 0x3b, 0xca, 0xef,
		0x2b, 0x81, 0x23, 0x00, 0x17, 0x74, 0x73, 0xa3, 0x3e, 0x7e, 0x81, 0xfe,
		0xbb, 0xf5, 0x3e, 0x55, 0x68, 0x22, 0x18, 0x14, 0x8e, 0x56, 0x0a, 0x61,
		0x21, 0x20, 0x87, 0x02, 0xf2, 0xa8, 0x9a, 0xe8, 0x7b, 0x41, 0xf1, 0x56,
		0xfe, 0x5f, 0x16, 0x56, 0x73, 0xfa, 0x3e, 0xa2, 0xa8, 0xac, 0x73, 0x0e,
		0x3b, 0x79, 0xea, 0x92, 0x7a, 0x1d, 0xf6, 0x88, 0xae, 0x07, 0xdd, 0x43,
		0xa0, 0x3d, 0x54, 0x6b, 0xae, 0x7b, 0xfd, 0x1d, 0x69, 0xdf, 0xd6, 0xd7,
		0xb1, 0x32, 0x1f, 0xa8, 0xa9, 0x5f, 0x4f, 0xbf, 0x57, 0x4a, 0x07, 0xf1,
		0xf8, 0x98, 0x60, 0xf7, 0xe6, 0x6c, 0x9b, 0xba, 0xba, 0x69, 0x49, 0xc1,
		0xe4, 0x5c, 0xf6, 0x0b, 0xc2, 0x39, 0x6a, 0xe8, 0xd1, 0xc6, 0x9a, 0x57,
		0x74, 0xb0, 0xce, 0x80, 0x52, 0x0c, 0x8a, 0xc1, 0xbf, 0x01, 0x00, 0x00,
		0xff, 0xff, 0xf4, 0x31, 0xf6, 0xbb, 0x02, 0x13, 0x00, 0x00,
	},
		"om_cluster_docs/replica_set.json",
	)
}

func om_cluster_docs_sharded_set_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xe4, 0x5a,
		0x4d, 0x73, 0xdb, 0x36, 0x13, 0xbe, 0xbf, 0x33, 0xef, 0x7f, 0xe0, 0xb0,
		0x57, 0xeb, 0x83, 0x94, 0xe2, 0xd8, 0x3e, 0xc5, 0x89, 0x0f, 0x39, 0xd4,
		0x75, 0xc6, 0x4e, 0xe2, 0x43, 0xe2, 0xd1, 0x80, 0x04, 0x48, 0xa2, 0x01,
		0x01, 0x0e, 0x00, 0xc9, 0x96, 0x35, 0xfc, 0xef, 0x05, 0x40, 0x8a, 0xdf,
		0xa4, 0x9c, 0xd6, 0x8d, 0x65, 0x55, 0x93, 0x43, 0x08, 0x3e, 0xc0, 0xee,
		0x3e, 0xbb, 0xfb, 0x40, 0x02, 0xbc, 0xf9, 0xff, 0xff, 0x2c, 0xcb, 0xb2,
		0x59, 0x22, 0x31, 0xa3, 0xc2, 0x3e, 0xb3, 0xb2, 0x01, 0xf3, 0xb1, 0x21,
		0xbb, 0xa7, 0x84, 0x01, 0xf8, 0x1e, 0x08, 0xa4, 0x5e, 0xd9, 0x93, 0x15,
		0xe0, 0x13, 0x82, 0xbd, 0x49, 0xcc, 0x68, 0xc8, 0xa0, 0x37, 0x8a, 0x63,
		0x31, 0x02, 0x4b, 0xc9, 0x62, 0xa0, 0xa7, 0xdb, 0x47, 0x3d, 0x73, 0x6f,
		0x31, 0x55, 0x8f, 0x7a, 0x75, 0xfb, 0xc3, 0xd9, 0xf7, 0xef, 0x3d, 0xb3,
		0xb3, 0xc9, 0x69, 0xbe, 0x88, 0x6d, 0x50, 0x17, 0xde, 0x57, 0xc4, 0x45,
		0xee, 0xda, 0xb7, 0xaa, 0x6f, 0xde, 0x12, 0x13, 0x68, 0x46, 0xcb, 0xc1,
		0xed, 0x67, 0xd3, 0x1e, 0xca, 0xe6, 0x60, 0xa9, 0x67, 0x1c, 0xcf, 0x8f,
		0x7a, 0x00, 0x01, 0x01, 0x2b, 0xc6, 0xb5, 0x9f, 0x76, 0x1f, 0x24, 0xc4,
		0x32, 0x77, 0x49, 0xc3, 0xe6, 0xee, 0xfc, 0xd4, 0x77, 0xa0, 0xeb, 0xbd,
		0x39, 0x3d, 0x3d, 0x45, 0x9e, 0x17, 0x38, 0x01, 0x0c, 0x66, 0x9e, 0x3f,
		0x45, 0x53, 0x34, 0xf3, 0x66, 0x41, 0xf0, 0xc6, 0x9f, 0x02, 0x10, 0xb8,
		0xbd, 0xcb, 0xc5, 0xe0, 0xe1, 0x4a, 0x54, 0x16, 0xec, 0x07, 0x62, 0xfa,
		0x44, 0x20, 0x83, 0x4b, 0x82, 0x0c, 0x33, 0x77, 0x7d, 0x98, 0x84, 0x00,
		0x19, 0x30, 0x1e, 0xeb, 0x85, 0x98, 0x78, 0xe8, 0x5d, 0x6b, 0xc9, 0x89,
		0x86, 0x44, 0x52, 0x26, 0xe2, 0x6c, 0x32, 0x09, 0x80, 0x90, 0x90, 0x8c,
		0xf3, 0xfc, 0x8d, 0x19, 0x0f, 0x55, 0x35, 0xd0, 0xe5, 0x43, 0x51, 0x0f,
		0xe6, 0x69, 0xf4, 0x70, 0x72, 0xbc, 0x38, 0x9e, 0x8f, 0x40, 0x0c, 0x1e,
		0x19, 0x1d, 0xcd, 0xc6, 0xee, 0xf8, 0xed, 0x58, 0x86, 0x8f, 0xbd, 0x56,
		0xee, 0x31, 0x75, 0xa7, 0xd3, 0x93, 0x84, 0x2c, 0xb5, 0xd7, 0x01, 0x20,
		0x02, 0x0d, 0x40, 0xbf, 0x7e, 0xb8, 0x46, 0x10, 0x0b, 0x79, 0x41, 0xc8,
		0x20, 0x0f, 0x15, 0xe8, 0x55, 0x51, 0xdd, 0xfd, 0x94, 0x54, 0xe0, 0x5f,
		0xf8, 0x93, 0x57, 0xae, 0x26, 0xa4, 0x8d, 0x4e, 0xeb, 0x43, 0x55, 0xdb,
		0x36, 0x05, 0xb1, 0xe9, 0x29, 0x43, 0xcf, 0xb6, 0xf6, 0xb7, 0x08, 0xdb,
		0x03, 0xfe, 0x8f, 0x65, 0x52, 0xad, 0xfd, 0xec, 0x85, 0x06, 0x94, 0xfd,
		0x81, 0x25, 0xe3, 0x98, 0x86, 0x6d, 0x58, 0xb5, 0x51, 0x22, 0x26, 0xe4,
		0xd6, 0xd8, 0x66, 0x43, 0x19, 0x44, 0x62, 0xfc, 0x6d, 0x7a, 0x97, 0xa6,
		0xb5, 0x76, 0x25, 0x2c, 0xfc, 0x04, 0x64, 0x54, 0x76, 0x39, 0x0b, 0x7b,
		0xba, 0x7c, 0x52, 0x1a, 0x1e, 0x81, 0x10, 0x51, 0x39, 0x56, 0xd8, 0xe6,
		0x5a, 0xd7, 0x4c, 0x02, 0x89, 0xea, 0x72, 0x62, 0xde, 0x09, 0xfc, 0x88,
		0x3e, 0x47, 0x1c, 0x89, 0x88, 0x11, 0x78, 0xf9, 0x5e, 0x21, 0x9c, 0xe9,
		0x74, 0xda, 0x60, 0xda, 0x96, 0x38, 0x2e, 0x51, 0x1f, 0xb9, 0x8e, 0xcc,
		0x9d, 0x97, 0x98, 0x9c, 0xd7, 0xb4, 0xe4, 0xc4, 0xcc, 0x4a, 0x38, 0xf3,
		0x91, 0x10, 0x59, 0xed, 0x6f, 0x36, 0xbf, 0x21, 0xe0, 0x47, 0x96, 0x09,
		0x38, 0x2d, 0x32, 0x51, 0x65, 0x06, 0xf0, 0x50, 0xb8, 0x8b, 0xe3, 0x0e,
		0x2f, 0x29, 0x92, 0xed, 0xd1, 0xcc, 0x06, 0xe3, 0xfa, 0x95, 0x7b, 0xa2,
		0xbc, 0xae, 0xbf, 0x4e, 0x9b, 0x41, 0x70, 0x94, 0x10, 0xec, 0x67, 0xda,
		0xd6, 0xbd, 0x98, 0x46, 0xdc, 0x20, 0xf9, 0x47, 0x9e, 0x1d, 0x11, 0x01,
		0x0e, 0x17, 0x9b, 0x0d, 0xc4, 0x2b, 0xeb, 0x9d, 0x92, 0x4c, 0xf4, 0x60,
		0xcd, 0x2c, 0x95, 0xa7, 0x1d, 0x76, 0x84, 0x4a, 0x86, 0x4a, 0x44, 0x8f,
		0x0d, 0xe8, 0xd5, 0xf2, 0xba, 0xf2, 0x41, 0x32, 0xd1, 0x33, 0x50, 0x91,
		0x5e, 0x08, 0x24, 0xd8, 0x69, 0x63, 0x2d, 0x24, 0x8a, 0x7f, 0x57, 0x89,
		0xee, 0xb1, 0x82, 0x84, 0xc4, 0x74, 0x1b, 0xab, 0x1d, 0x60, 0x82, 0xba,
		0xda, 0xc7, 0x4e, 0xda, 0xbe, 0xac, 0x45, 0xa5, 0xd8, 0x16, 0x3a, 0x5d,
		0x5b, 0xcf, 0x4c, 0x61, 0x35, 0x1c, 0xab, 0x14, 0xc1, 0x51, 0x6f, 0x91,
		0xcb, 0x08, 0x8b, 0x76, 0x7d, 0xff, 0xcb, 0x35, 0xd9, 0xd5, 0xdc, 0x9d,
		0xae, 0xe4, 0x65, 0xfa, 0x79, 0x9d, 0x18, 0x4c, 0x16, 0x6c, 0x0d, 0xb1,
		0x2a, 0x55, 0x25, 0x93, 0x87, 0xea, 0x4b, 0xd5, 0x89, 0xd1, 0x8d, 0x1f,
		0xa1, 0x18, 0x94, 0xe2, 0xf3, 0x26, 0x6f, 0x08, 0x55, 0xf4, 0x38, 0xb0,
		0xde, 0x11, 0xa5, 0xd3, 0xa9, 0x7a, 0x40, 0x4a, 0x49, 0xd3, 0xf4, 0x68,
		0xb3, 0x99, 0xe0, 0x60, 0xdb, 0x02, 0xea, 0x41, 0xf7, 0x45, 0xda, 0x6c,
		0x9e, 0xbc, 0x5a, 0x55, 0x39, 0xe6, 0x3a, 0xd2, 0x95, 0xe6, 0x05, 0x86,
		0x65, 0xa1, 0x4e, 0x3b, 0x33, 0x1c, 0xa3, 0xd8, 0x43, 0xbc, 0x67, 0x63,
		0x36, 0x0e, 0x74, 0x0f, 0x57, 0xd6, 0x6f, 0x52, 0x5f, 0x83, 0x00, 0xae,
		0x76, 0x71, 0xc4, 0xaf, 0x28, 0x59, 0xef, 0xd8, 0x2c, 0x0c, 0x3c, 0xc2,
		0x10, 0x22, 0xfa, 0x24, 0xa4, 0x2a, 0xa1, 0x22, 0x38, 0x04, 0x47, 0xbe,
		0xda, 0x8e, 0x94, 0xa1, 0x45, 0x1e, 0xac, 0xd2, 0x89, 0xa1, 0xd9, 0x09,
		0xc7, 0x4a, 0x12, 0xa5, 0xf6, 0xc9, 0x19, 0xc2, 0x09, 0xf5, 0x0d, 0x03,
		0x5d, 0x20, 0x02, 0xd6, 0xbb, 0x02, 0x5d, 0x31, 0x69, 0x94, 0xcc, 0xe9,
		0xc6, 0x34, 0x5b, 0x74, 0xfb, 0xd9, 0x4d, 0xef, 0xa0, 0x7f, 0x2f, 0x46,
		0xef, 0xdb, 0x43, 0xa1, 0xd7, 0xdd, 0x4b, 0x7a, 0x4f, 0xf6, 0x8c, 0xde,
		0xf6, 0xf0, 0xdd, 0xf0, 0x16, 0xb4, 0x53, 0x8e, 0x9c, 0xff, 0x92, 0x1c,
		0x39, 0x8b, 0xd3, 0x3d, 0x4b, 0xe8, 0x41, 0xc9, 0x91, 0xb3, 0x70, 0x3a,
		0x77, 0xb7, 0x62, 0xfa, 0x2b, 0xe2, 0x77, 0x1f, 0xf5, 0x48, 0xf1, 0xdb,
		0xd9, 0xae, 0xc5, 0xf4, 0x43, 0x10, 0xa4, 0xce, 0x1f, 0xfd, 0x87, 0x2a,
		0x48, 0xee, 0xc2, 0xe9, 0x3d, 0xe4, 0x30, 0xd3, 0x5f, 0x51, 0xc7, 0xec,
		0xa3, 0x22, 0x29, 0x7e, 0x67, 0x87, 0xc2, 0xef, 0x3e, 0x2a, 0x92, 0xe2,
		0x77, 0xbe, 0x67, 0xfc, 0x3e, 0xa7, 0x22, 0xf9, 0x8c, 0x06, 0x38, 0xbc,
		0x16, 0x87, 0x29, 0x49, 0xdb, 0xe8, 0xba, 0x7f, 0x92, 0x16, 0xf0, 0x57,
		0xd4, 0x22, 0x2f, 0x2b, 0x41, 0x05, 0x9f, 0xfb, 0xb6, 0x49, 0xbf, 0x52,
		0xc9, 0x29, 0xf8, 0xdc, 0xb7, 0x2d, 0x72, 0xb7, 0xc4, 0x64, 0x8f, 0xc5,
		0x29, 0x91, 0x11, 0x4b, 0x4c, 0xc3, 0xba, 0x56, 0x74, 0x09, 0x8f, 0x41,
		0xfe, 0x4d, 0x4d, 0x91, 0x20, 0x1c, 0x3c, 0xac, 0x37, 0xa0, 0xdd, 0x67,
		0x51, 0x05, 0xd4, 0x68, 0x5b, 0x81, 0x7c, 0xbe, 0xda, 0xfa, 0x79, 0x3f,
		0x87, 0x1b, 0xaa, 0xea, 0xa7, 0xf3, 0xa2, 0x7e, 0x0e, 0x17, 0x6a, 0xd5,
		0x4f, 0xb7, 0xcf, 0xcf, 0x8e, 0xda, 0xea, 0xda, 0x7c, 0x68, 0xf5, 0xf0,
		0xbb, 0xdc, 0x96, 0x3b, 0x37, 0xaa, 0xac, 0x8f, 0x6e, 0x10, 0x5f, 0x21,
		0xde, 0x17, 0x4f, 0x0d, 0x74, 0x9d, 0x1d, 0x6a, 0xee, 0xdc, 0x01, 0x7d,
		0x46, 0x08, 0xf2, 0xcb, 0x4b, 0xa2, 0xde, 0x26, 0xc8, 0xdb, 0x40, 0x9f,
		0xc1, 0xda, 0x67, 0x25, 0xe5, 0x36, 0xc4, 0x02, 0x78, 0x04, 0x29, 0x06,
		0x1b, 0xa2, 0xa0, 0xa1, 0xec, 0xd3, 0x3d, 0xcc, 0x4e, 0x83, 0xad, 0x04,
		0x08, 0x71, 0xcf, 0x38, 0xb4, 0x6a, 0x67, 0xc2, 0x06, 0xf4, 0x45, 0xe8,
		0xa0, 0xec, 0xde, 0x2b, 0x55, 0x1b, 0xaa, 0x60, 0xd8, 0x3a, 0x46, 0x54,
		0x9e, 0x2b, 0xf3, 0x97, 0xc8, 0x8f, 0x00, 0xc5, 0x22, 0x6e, 0xf7, 0x98,
		0x7d, 0xf3, 0xe1, 0xfa, 0xfc, 0x72, 0x74, 0xf3, 0xf1, 0x7c, 0x54, 0xa9,
		0xa2, 0x0a, 0x59, 0xf6, 0x0f, 0xa4, 0xa4, 0x43, 0xfb, 0xa3, 0xfe, 0xd3,
		0x70, 0x45, 0x8d, 0x98, 0xd3, 0xfb, 0xb3, 0xca, 0x19, 0xfd, 0x9f, 0xcc,
		0x13, 0xb5, 0x03, 0xfa, 0x8c, 0xcc, 0x6c, 0x68, 0xc1, 0xe2, 0xb1, 0x5e,
		0xaf, 0xb2, 0xc6, 0x52, 0x85, 0x22, 0x6e, 0x01, 0x95, 0x9a, 0x90, 0x9a,
		0x6f, 0xcd, 0x2a, 0xb5, 0xa1, 0xa7, 0x2c, 0x01, 0x18, 0x63, 0xda, 0x4a,
		0x8c, 0x8d, 0x29, 0x96, 0x86, 0xba, 0x7e, 0xe6, 0x72, 0x24, 0x67, 0xe6,
		0xc6, 0xf3, 0xc9, 0x57, 0xc1, 0x43, 0x66, 0x2b, 0x4b, 0x2a, 0x4c, 0x5e,
		0x8e, 0x97, 0xd9, 0x9d, 0xd7, 0xee, 0x3b, 0xbe, 0x76, 0x91, 0x1b, 0x36,
		0xf2, 0xc4, 0x36, 0xaf, 0xce, 0x6a, 0xeb, 0xd5, 0x3b, 0xfb, 0xb5, 0x32,
		0x75, 0x6e, 0xa0, 0x1d, 0x3c, 0x75, 0xcc, 0xfe, 0xa7, 0x26, 0x39, 0x02,
		0xf0, 0x9c, 0xae, 0x2f, 0x80, 0x04, 0x9e, 0xfe, 0x13, 0x85, 0x5f, 0x63,
		0x55, 0xe7, 0xd3, 0x44, 0xf9, 0xbc, 0xa6, 0x09, 0xf3, 0x01, 0x79, 0x42,
		0xc0, 0xb7, 0xea, 0x2b, 0xc2, 0xaf, 0x0a, 0x75, 0xd0, 0xde, 0xcf, 0x14,
		0x7e, 0x76, 0x9f, 0xdd, 0x55, 0xf4, 0x9d, 0xfa, 0x64, 0x04, 0x44, 0x7d,
		0xc3, 0x41, 0x99, 0x82, 0xdc, 0x35, 0xa4, 0xb2, 0xa6, 0x7f, 0x5a, 0x59,
		0xdb, 0x82, 0xa7, 0xd6, 0x55, 0xff, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
		0x94, 0x3d, 0xd4, 0x32, 0xde, 0x22, 0x00, 0x00,
	},
		"om_cluster_docs/sharded_set.json",
	)
}

func om_cluster_docs_standalone_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x56,
		0xcf, 0x73, 0xda, 0x38, 0x14, 0xbe, 0xf3, 0x57, 0x30, 0x3e, 0x17, 0x70,
		0x08, 0xc9, 0x26, 0xb9, 0xd1, 0xe4, 0xd0, 0xc3, 0x66, 0xda, 0x49, 0xda,
		0xe6, 0xd0, 0x76, 0x98, 0x67, 0x4b, 0xb6, 0xb5, 0x95, 0x25, 0x8f, 0x24,
		0x92, 0x90, 0x0c, 0xff, 0xfb, 0x4a, 0xb2, 0x01, 0x0b, 0x4b, 0xd4, 0x9d,
		0x29, 0x9b, 0xf5, 0x09, 0x9e, 0x3e, 0xbd, 0x5f, 0xfa, 0xbe, 0x27, 0xbd,
		0x0e, 0x86, 0xfa, 0x8b, 0x78, 0xa5, 0x08, 0x67, 0x32, 0xba, 0x1a, 0xd6,
		0x06, 0x6b, 0x44, 0xfc, 0x89, 0x51, 0x0e, 0xe8, 0x3d, 0x48, 0xac, 0x57,
		0xa2, 0xc9, 0x23, 0x88, 0x09, 0x25, 0xc9, 0xa4, 0xe4, 0x2c, 0xe7, 0x28,
		0x19, 0x95, 0xa5, 0x1c, 0xc1, 0x52, 0xf1, 0x12, 0xcc, 0xee, 0xe8, 0x9d,
		0x7f, 0xeb, 0x03, 0x61, 0xfa, 0xaf, 0xf1, 0x1d, 0x5d, 0x5f, 0x7d, 0xff,
		0x1e, 0xd8, 0x6c, 0xf7, 0xae, 0x6b, 0x17, 0x91, 0xc5, 0xdc, 0x24, 0x5f,
		0xb1, 0x90, 0x4d, 0x5a, 0xdf, 0x5a, 0x79, 0x25, 0x4b, 0x42, 0x91, 0x35,
		0x6e, 0x6d, 0x9b, 0xef, 0xb5, 0x63, 0xa9, 0x77, 0x10, 0x65, 0xf0, 0xe7,
		0xb3, 0x77, 0xfe, 0xf5, 0x8c, 0xc2, 0x23, 0x17, 0x26, 0xc3, 0x28, 0x80,
		0xc8, 0x89, 0x6a, 0xb2, 0x31, 0xa8, 0xd9, 0x74, 0x76, 0x99, 0x9e, 0xa0,
		0x69, 0x72, 0x76, 0x79, 0x79, 0x89, 0x93, 0x24, 0x3b, 0xc9, 0x50, 0x76,
		0x9a, 0xa4, 0x31, 0x8e, 0xf1, 0x69, 0x72, 0x9a, 0x65, 0x67, 0x69, 0x0c,
		0x90, 0x4d, 0x43, 0xde, 0x4a, 0x78, 0xfe, 0x28, 0x5b, 0xfe, 0x82, 0x38,
		0xc2, 0xfa, 0xe1, 0x38, 0x5a, 0x52, 0x6c, 0x5b, 0xf2, 0x23, 0x00, 0xa9,
		0x28, 0xa8, 0x8c, 0x8b, 0xd2, 0xb8, 0xe1, 0xf2, 0x39, 0xe4, 0x69, 0x29,
		0xa8, 0x41, 0x14, 0x4a, 0x55, 0xf2, 0x6a, 0x32, 0xc9, 0x40, 0x2a, 0x44,
		0xc7, 0xcd, 0x99, 0x8d, 0xb9, 0xc8, 0x35, 0x01, 0xd8, 0xf2, 0x79, 0x4b,
		0x01, 0xfb, 0x6f, 0xf4, 0x7c, 0x71, 0xbe, 0x38, 0x9f, 0x8d, 0xa0, 0x84,
		0x17, 0xce, 0x46, 0xa7, 0xe3, 0xe9, 0xf8, 0xaf, 0xb1, 0xca, 0x5f, 0x42,
		0x41, 0x9e, 0x08, 0x9b, 0xc6, 0xf1, 0x45, 0x45, 0x97, 0x26, 0xe5, 0x0c,
		0xa8, 0xc4, 0x61, 0xe4, 0xd7, 0xeb, 0x3b, 0x8c, 0x88, 0x54, 0x37, 0x94,
		0x1e, 0x6a, 0x41, 0x0b, 0xf9, 0x71, 0x4b, 0xe6, 0x60, 0x37, 0x5a, 0xe8,
		0x2f, 0xa2, 0xaf, 0xdf, 0xf6, 0x41, 0x74, 0xc0, 0x6b, 0xc7, 0xd2, 0x8a,
		0x1b, 0x31, 0x28, 0xad, 0x7c, 0x6c, 0x5b, 0x1a, 0x9e, 0x37, 0xeb, 0x51,
		0x02, 0xe9, 0xcf, 0x65, 0x15, 0xe0, 0x79, 0xc1, 0xa5, 0xda, 0x6c, 0x7e,
		0x7d, 0x65, 0x1c, 0x61, 0x39, 0xfe, 0x16, 0xff, 0x58, 0xaf, 0xdb, 0x4a,
		0xa3, 0x3c, 0xff, 0x04, 0xaa, 0xd8, 0xe9, 0x93, 0xe7, 0x01, 0x7d, 0x4e,
		0xea, 0x60, 0x23, 0xc8, 0x31, 0x53, 0x63, 0x8d, 0xdb, 0x73, 0x73, 0xc7,
		0x15, 0x28, 0xec, 0x8c, 0x00, 0xbb, 0x24, 0xc9, 0x0b, 0xfe, 0x5c, 0x08,
		0x2c, 0x0b, 0x4e, 0xd1, 0xed, 0x7b, 0x0d, 0x38, 0x89, 0xe3, 0xd8, 0x6d,
		0x57, 0xa4, 0x48, 0xb9, 0x03, 0x7d, 0x10, 0xa6, 0x94, 0xe9, 0x6c, 0xe0,
		0x36, 0xc7, 0x54, 0xbd, 0xd5, 0x37, 0x51, 0x5c, 0x10, 0x96, 0xff, 0x47,
		0xa5, 0xef, 0x02, 0xbe, 0x6d, 0xf9, 0x16, 0x5c, 0x09, 0x9e, 0x62, 0x29,
		0xf1, 0x7e, 0xd1, 0x20, 0x72, 0x39, 0x5d, 0x9c, 0x77, 0x53, 0x60, 0x58,
		0x75, 0x8c, 0xb5, 0x27, 0x2e, 0xcc, 0xca, 0xf4, 0x42, 0x67, 0xe4, 0xac,
		0xae, 0xf7, 0xf2, 0x93, 0xba, 0x78, 0x5d, 0xb8, 0xdf, 0x0b, 0x4a, 0x9c,
		0x36, 0x3e, 0xa6, 0x50, 0x4d, 0xcc, 0x06, 0xbc, 0xed, 0x26, 0x02, 0x05,
		0xd1, 0xe1, 0x00, 0x2b, 0xa9, 0x70, 0xf9, 0xb7, 0xee, 0xaa, 0x3f, 0x04,
		0x96, 0x8a, 0xb0, 0x7a, 0xd2, 0xeb, 0x38, 0x19, 0xa1, 0xd8, 0x23, 0xb8,
		0xa8, 0xea, 0xe6, 0xb1, 0x92, 0xad, 0x73, 0x5d, 0x18, 0x26, 0x6c, 0xb2,
		0xb2, 0x67, 0xe8, 0x26, 0x35, 0xf0, 0xa4, 0xd7, 0x9b, 0x4b, 0xc7, 0x24,
		0x80, 0x67, 0x1c, 0x04, 0x53, 0x69, 0xe8, 0xf1, 0x79, 0x55, 0x59, 0x5c,
		0x5d, 0x6d, 0x1b, 0xf0, 0xb8, 0x9b, 0x43, 0xf5, 0x4c, 0x69, 0xad, 0x69,
		0xce, 0x17, 0xf7, 0x69, 0x81, 0x4b, 0xd8, 0x4d, 0xab, 0x33, 0x97, 0x7e,
		0x02, 0x57, 0x94, 0xa4, 0x70, 0x8f, 0x55, 0x7b, 0x44, 0x46, 0x82, 0xbb,
		0x37, 0x48, 0x24, 0x0b, 0x10, 0x48, 0x2b, 0xa6, 0xb1, 0x0d, 0xb6, 0xfe,
		0xdd, 0x57, 0x82, 0x51, 0xd9, 0x17, 0x89, 0xed, 0xed, 0x19, 0x7e, 0x10,
		0x18, 0xeb, 0xa7, 0x27, 0x54, 0x17, 0x3e, 0xac, 0x40, 0xca, 0x27, 0x2e,
		0xd0, 0xd0, 0xad, 0x1c, 0xe9, 0xcc, 0xf8, 0xaa, 0xd4, 0xfa, 0x9c, 0xeb,
		0x30, 0xb7, 0x38, 0x2d, 0x80, 0x11, 0x59, 0x76, 0xaf, 0xfa, 0xe8, 0xfe,
		0xfa, 0x6e, 0x7e, 0x3b, 0xba, 0xff, 0x30, 0x1f, 0x9d, 0xec, 0x28, 0xd0,
		0x1e, 0xba, 0x3f, 0xf1, 0xaa, 0x89, 0xa5, 0x7f, 0xed, 0x85, 0xd1, 0x16,
		0x4b, 0x40, 0x87, 0x68, 0xff, 0xf0, 0x44, 0x3a, 0x2c, 0x4b, 0x39, 0xcb,
		0x48, 0x43, 0xbc, 0x05, 0x2f, 0xc7, 0xc6, 0x63, 0x3b, 0x57, 0x22, 0x21,
		0xa1, 0x18, 0x75, 0x2f, 0xb0, 0x68, 0xa9, 0xbb, 0x21, 0x6f, 0x30, 0xc5,
		0xca, 0x2e, 0xb7, 0x2f, 0xa1, 0x7a, 0xed, 0x01, 0x58, 0xb3, 0xe4, 0x54,
		0xe5, 0x95, 0xa6, 0xc9, 0x12, 0x90, 0x7e, 0x08, 0xf8, 0xf4, 0xb2, 0x3d,
		0x33, 0xef, 0xdd, 0xe5, 0x7f, 0x0c, 0xf5, 0xf0, 0xeb, 0xf8, 0x37, 0xc0,
		0x54, 0xdf, 0xd4, 0x0a, 0x8b, 0xdb, 0x7a, 0x8a, 0x76, 0x2f, 0x3f, 0xf3,
		0xad, 0x3b, 0x56, 0xcf, 0xe5, 0x6b, 0xeb, 0xdf, 0x30, 0x65, 0x7f, 0x28,
		0xfb, 0x2a, 0x24, 0x1a, 0xe2, 0xe7, 0xcd, 0xa1, 0x71, 0xf4, 0xbf, 0xef,
		0xe4, 0xdc, 0xe2, 0xfd, 0x7d, 0xf4, 0x7b, 0xf9, 0x73, 0x29, 0x08, 0x0c,
		0x68, 0xce, 0x56, 0x37, 0x7a, 0xa6, 0x27, 0xe6, 0x65, 0xff, 0x36, 0x59,
		0x18, 0x22, 0xd8, 0x2e, 0x1c, 0x2d, 0x15, 0xca, 0x53, 0xa0, 0x7d, 0x1b,
		0xf2, 0x20, 0x88, 0x7a, 0xab, 0x56, 0xfc, 0x2a, 0xfe, 0x6f, 0x0b, 0xab,
		0xfd, 0xd0, 0x3b, 0xa2, 0xa8, 0x9c, 0x3a, 0x87, 0x9d, 0x38, 0xdb, 0x94,
		0x82, 0x80, 0x03, 0xa2, 0x0b, 0x74, 0xb7, 0x4f, 0x6b, 0xfb, 0x6a, 0xcd,
		0x77, 0xae, 0x7f, 0x22, 0xec, 0xaf, 0xf5, 0x75, 0xac, 0xc8, 0x3d, 0x35,
		0xf5, 0xfb, 0xe1, 0x0f, 0x4a, 0xa9, 0x17, 0x8f, 0x8f, 0xd9, 0xec, 0x60,
		0xcc, 0x7d, 0x53, 0x57, 0x37, 0x7b, 0x52, 0xb0, 0x31, 0x17, 0x61, 0x41,
		0x78, 0x9f, 0x1a, 0xe6, 0x69, 0xe3, 0xbc, 0x57, 0x8c, 0xb3, 0xce, 0x03,
		0x65, 0x3d, 0x58, 0x0f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x44, 0xd4,
		0x65, 0x75, 0x6d, 0x11, 0x00, 0x00,
	},
		"om_cluster_docs/standalone.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"om_cluster_docs/replica_set.json": om_cluster_docs_replica_set_json,
	"om_cluster_docs/sharded_set.json": om_cluster_docs_sharded_set_json,
	"om_cluster_docs/standalone.json": om_cluster_docs_standalone_json,
}
