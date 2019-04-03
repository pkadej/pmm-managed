// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package agents

import (
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/agentpb"

	"github.com/percona/pmm-managed/models"
)

func nodeExporterConfig(node *models.Node, exporter *models.Agent) *agentpb.SetStateRequest_AgentProcess {
	tdp := templateDelimsPair(
		pointer.GetString(exporter.MetricsURL),
	)

	args := []string{
		// "--collector.ntp", disabled for now
		//"--collector.runit", disabled for now
		//"--collector.supervisord", disabled for now
		// "--collector.tcpstat", disabled for now

		// TODO
		// "--collector.textfile",
		// "--collector.textfile.directory",

		"--web.listen-address=:" + tdp.left + " .listen_port " + tdp.right,
	}

	// do not enable Linux-specific collectors on macOS, that's useful for our development
	if pointer.GetString(node.Distro) != "darwin" {
		args = append(args,
			// enable disabled by default
			"--collector.buddyinfo",
			"--collector.drbd",
			"--collector.interrupts",
			"--collector.ksmd",
			"--collector.logind",
			"--collector.meminfo_numa",
			"--collector.mountstats",
			"--collector.processes",
			"--collector.qdisc",
			"--collector.systemd",
			"--collector.wifi",
		)
	}

	if pointer.GetString(exporter.MetricsURL) != "" {
		args = append(args, "--web.telemetry-path="+*exporter.MetricsURL)
	}

	sort.Strings(args)

	return &agentpb.SetStateRequest_AgentProcess{
		Type:               agentpb.Type_NODE_EXPORTER,
		TemplateLeftDelim:  tdp.left,
		TemplateRightDelim: tdp.right,
		Args:               args,
	}
}
