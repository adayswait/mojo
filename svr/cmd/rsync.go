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
			mlog.Infof("rsync cmd : %s exec succeed", rsyncCmd)
			return nil
		} else {
			mlog.Errorf("rsync cmd %s exec failed, err:%s, match:%s, ret:%s",
				rsyncCmd, errRsync, matchedRsync, retRsync)
			return fmt.Errorf("rsync cmd:%s exec failed", rsyncCmd)
		}
	} else {
		mlog.Infof("rsync rsh failed, cmd:%s, ret:%s,matched:%s,err:%s",
			rsyncCmd, retSSH, matchedSSH, errSSH)
		return fmt.Errorf("rsync rsh failed, cmd:%s", rsyncCmd)
	}
	return nil
}
