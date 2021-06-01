package main

import (
	"os/user"

	jobspec "github.com/coreos/gangplank/spec"
	flag "github.com/spf13/pflag"
)

// Flags has the configuration flags.
var specCommonFlags = flag.NewFlagSet("", flag.ContinueOnError)

// sshFlags are specific to minio and remote podman
var (
	sshFlags = flag.NewFlagSet("", flag.ContinueOnError)

	// minioSshRemoteHost is the SSH remote host to forward the local
	// minio instance over.
	minioSshRemoteHost string

	// minioSshRemoteUser is the name of the SSH user to use with minioSshRemoteHost
	minioSshRemoteUser string

	// minioSshRemoteKey is the SSH key to use with minioSshRemoteHost
	minioSshRemoteKey string
)

// cosaKolaTests are used to generate automatic Kola stages.
var cosaKolaTests []string

func init() {
	specCommonFlags.StringSliceVar(&generateCommands, "singleCmd", []string{}, "commands to run in stage")
	specCommonFlags.StringSliceVar(&generateSingleRequires, "singleReq", []string{}, "artifacts to require")
	specCommonFlags.StringVarP(&cosaSrvDir, "srvDir", "S", "", "directory for /srv; in pod mount this will be bind mounted")
	jobspec.AddKolaTestFlags(&cosaKolaTests, specCommonFlags)

	user, _ := user.Current()
	sshFlags.StringVar(&minioSshRemoteHost, "forwardMinioSSH", containerHost(), "forward and use minio to ssh host")
	sshFlags.StringVar(&minioSshRemoteUser, "sshUser", user.Username, "name of SSH; used with forwardMinioSSH")
	sshFlags.StringVar(&minioSshRemoteKey, "sshKey", "", "path to SSH key; used with forwardMinioSSH")
}