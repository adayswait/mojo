package cmd

import (
	"fmt"
	"github.com/adayswait/mojo/mlog"
	"github.com/google/goexpect"
	"regexp"
	"time"
)

//rsync localPath to remotePath using ssh
func Rsync(
	localPath, remotePath string,
	remoteUser, remotePasswd string,
	remoteIp, remotePort string,
	options []string, timeout time.Duration,
) error {
	rshCmd := fmt.Sprintf("--rsh=ssh -p %s", remotePort)
	remoteServer := fmt.Sprintf("%s@%s:%s", remoteUser, remoteIp, remotePath)
	rsyncCmd := []string{"rsync", rshCmd, localPath, remoteServer}
	rsyncCmd = append(rsyncCmd, options...)
	expectRsync, _, _ := expect.SpawnWithArgs(rsyncCmd, -1)
	expectPasswd := regexp.MustCompile("password:")
	retSSH, matchedSSH, errSSH :=
		expectRsync.Expect(expectPasswd, timeout)
	if len(matchedSSH) == 1 && errSSH == nil {
		expectRsync.Send(remotePasswd + "\n")
		retRsync, matchedRsync, errRsync := expectRsync.Expect(
			regexp.MustCompile("speedup is"), timeout)
		if errRsync == nil && len(matchedRsync) == 1 {
			mlog.Log("rsync succeed", rsyncCmd)
			return nil
		} else {
			mlog.Log("rsync failed", rsyncCmd, errRsync, matchedRsync, retRsync)
			return fmt.Errorf("rsync cmd:%s failed", rsyncCmd)
		}
	} else {
		mlog.Log("rsync rsh failed", rsyncCmd, retSSH, matchedSSH, errSSH)
		return fmt.Errorf("rsync-ssh cmd:%s failed", rsyncCmd)
	}
	return nil
}
