//go:build netbsd || plan9 || solaris
// +build netbsd plan9 solaris

package stats

import "github.com/seaweedfs/seaweedfs/weed/pb/volume_server_pb"

func fillInDiskStatus(status *volume_server_pb.DiskStatus) {
	return
}
