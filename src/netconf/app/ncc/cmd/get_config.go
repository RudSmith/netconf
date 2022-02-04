// -*- coding: utf-8 -*-

// Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ncccmd

import (
	nc "github.com/juniper/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GetConfigCommand struct {
	Command
	Output string
}

func (c *GetConfigCommand) SetFlags(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "-", "Output filename.")
	return c.Command.SetFlags(cmd)
}

func (c *GetConfigCommand) GetConfig(source string) error {
	c.Command.Init()

	client, err := c.Client()
	if err != nil {
		log.Errorf("Client initialize error. %s", err)
		return err
	}
	defer client.Close()

	reply, err := client.Exec(nc.MethodGetConfig(source))
	if err != nil {
		log.Errorf("Exec error. %s", err)
		return err
	}

	return outputRPCReply(c.Output, reply)
}

func GetConfigCmd() *cobra.Command {
	get := GetConfigCommand{}
	c := get.SetFlags(&cobra.Command{
		Use:   "get-config [running/startup/candidate]",
		Short: "get-config command.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.GetConfig(args[0])
		},
	})

	return c
}
