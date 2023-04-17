module gitee.com/chunanyong/zorm-examples

go 1.13

require (
	// sql/gokb v0.0.0
	kingbase.com/gokb v0.0.0
	gitee.com/chunanyong/dm v1.8.11
	gitee.com/chunanyong/zorm v1.6.7
	github.com/cectc/hptx v1.0.5
	github.com/go-sql-driver/mysql v1.7.0
)

replace (
	kingbase.com/gokb => ./gokb
// kingbase.com/gokb => /Users/hjm/kingbase.com/gokb
)
