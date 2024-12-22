package aliyunoss

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/grafana/tempo/pkg/util"
)

type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

type HTTPMaxConns struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int
}

type Config struct {
	Endpoint            string                  `yaml:"endpoint"`             // OSS endpoint
	AccessKeyID         string                  `yaml:"access_key_id"`        // AccessId
	AccessKeySecret     string                  `yaml:"access_key_secret"`    // AccessKey
	RetryTimes          uint                    `yaml:"retry_times"`          // Retry count by default it's 5.
	UserAgent           string                  `yaml:"user_agent"`           // SDK name/version/system information
	IsDebug             bool                    `yaml:"is_debug"`             // Enable debug mode. Default is false.
	Timeout             uint                    `yaml:"timeout"`              // Timeout in seconds. By default it's 60.
	SecurityToken       string                  `yaml:"security_token"`       // STS Token
	IsCname             bool                    `yaml:"is_cname"`             // If cname is in the endpoint.
	IsPathStyle         bool                    `yaml:"is_path_style"`        // If Path Style is in the endpoint.
	HTTPTimeout         HTTPTimeout             `yaml:"http_timeout"`         // HTTP timeout
	HTTPMaxConns        HTTPMaxConns            `yaml:"http_max_conns"`       // Http max connections
	IsUseProxy          bool                    `yaml:"is_use_proxy"`         // Flag of using proxy.
	ProxyHost           string                  `yaml:"proxy_host"`           // Flag of using proxy host.
	IsAuthProxy         bool                    `yaml:"is_auth_proxy"`        // Flag of needing authentication.
	ProxyUser           string                  `yaml:"proxy_user"`           // Proxy user
	ProxyPassword       string                  `yaml:"proxy_password"`       // Proxy password
	IsEnableMD5         bool                    `yaml:"is_enable_md_5"`       // Flag of enabling MD5 for upload.
	MD5Threshold        int64                   `yaml:"md_5_threshold"`       // Memory footprint threshold for each MD5 computation (16MB is the default), in byte. When the data is more than that, temp file is used.
	IsEnableCRC         bool                    `yaml:"is_enable_crc"`        // Flag of enabling CRC for upload.
	LogLevel            int                     `yaml:"log_level"`            // Log level
	Logger              *log.Logger             `yaml:"logger"`               // For write log
	UploadLimitSpeed    int                     `yaml:"upload_limit_speed"`   // Upload limit speed:KB/s, 0 is unlimited
	UploadLimiter       *oss.OssLimiter         `yaml:"upload_limiter"`       // Bandwidth limit reader for upload
	DownloadLimitSpeed  int                     `yaml:"download_limit_speed"` // Download limit speed:KB/s, 0 is unlimited
	DownloadLimiter     *oss.OssLimiter         `yaml:"download_limiter"`     // Bandwidth limit reader for download
	CredentialsProvider oss.CredentialsProvider `yaml:"credentials_provider"` // User provides interface to get AccessKeyID, AccessKeySecret, SecurityToken
	LocalAddr           net.Addr                `yaml:"local_addr"`           // local client host info
	UserSetUa           bool                    `yaml:"user_set_ua"`          // UserAgent is set by user or not
	AuthVersion         oss.AuthVersionType     `yaml:"auth_version"`         //  v1 or v2, v4 signature,default is v1
	AdditionalHeaders   []string                `yaml:"additional_headers"`   //  special http headers needed to be sign
	RedirectEnabled     bool                    `yaml:"redirect_enabled"`     //  only effective from go1.7 onward, enable http redirect or not
	InsecureSkipVerify  bool                    `yaml:"insecure_skip_verify"` //  for https, Whether to skip verifying the server certificate file
	Region              string                  `yaml:"region"`               //  such as cn-hangzhou
	CloudBoxId          string                  `yaml:"cloud_box_id"`         //
	Product             string                  `yaml:"product"`              //  oss or oss-cloudbox, default is oss
	VerifyObjectStrict  bool                    `yaml:"verify_object_strict"` //  a flag of verifying object name strictly. Default is enable.
}

func (cfg *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	f.StringVar(&cfg.Bucket, util.PrefixConfig(prefix, "aliyunoss.bucket"), "", "s3aliyunoss bucket to store blocks in.")
	f.StringVar(&cfg.Endpoint, util.PrefixConfig(prefix, "aliyunoss.endpoint"), "", "aliyunoss endpoint to push blocks to.")
	f.StringVar(&cfg.AccessKeySecret, util.PrefixConfig(prefix, "aliyunoss.access_key_secret"), "", "aliyunoss access key secret.")
	f.StringVar(&cfg.AccessKeyID, util.PrefixConfig(prefix, "aliyunoss.access_key_id"), "", "aliyunoss access key id.")
}

func (cfg *Config) PathMatches(other *Config) bool {
	// S3 bucket names are globally unique
	return cfg.Bucket == other.Bucket
}
