rm go.mod
go mod init github.com/purpose168/GoAdmin-example
cat >> go.mod << EOF

replace (
	github.com/purpose168/GoAdmin => github.com/purpose168/GoAdmin v0.0.0-20260104141321-fcc00eb84719
	github.com/purpose168/GoAdmin-themes => github.com/purpose168/GoAdmin-themes v0.0.0-20260104133356-8e29cafd3a6d
)
EOF
go mod tidy