// replication-manager - Replication Manager Monitoring and CLI for MariaDB
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <stephane@mariadb.com>
// This source code is licensed under the GNU General Public License, version 3.

package cluster

import "time"

func (cluster *Cluster) testSlaReplAllSlavesDelayNoSemiSync(conf string, test string) bool {
	if cluster.initTestCluster(conf, test) == false {
		return false
	}

	cluster.conf.MaxDelay = 2
	err := cluster.disableSemisync()
	if err != nil {
		cluster.LogPrintf("ERROR : %s", err)
		cluster.closeTestCluster(conf, test)
		return false
	}

	cluster.sme.ResetUpTime()
	time.Sleep(3 * time.Second)
	sla1 := cluster.sme.GetUptimeFailable()
	cluster.DelayAllSlaves()
	sla2 := cluster.sme.GetUptimeFailable()
	err = cluster.enableSemisync()
	if err != nil {
		cluster.LogPrintf("ERROR : %s", err)
		cluster.closeTestCluster(conf, test)
		return false
	}
	if sla2 == sla1 {
		cluster.closeTestCluster(conf, test)
		return false
	} else {
		cluster.closeTestCluster(conf, test)
		return true
	}
}
