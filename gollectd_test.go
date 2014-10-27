package gollectd

import (
	//"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FormatTests struct {
	Description string
	Expected    string
	Packet      Packet
}

var formatTests = []FormatTests{
	{
		Description: "leeloo/cpu-0/cpu-idle",
		Expected:    "leeloo/cpu-0/cpu-idle",
		Packet: Packet{
			Hostname:       "leeloo",
			Plugin:         "cpu",
			PluginInstance: "0",
			Type:           "cpu",
			TypeInstance:   "idle",
		},
	},
	{
		Description: "alyja/memory/memory-used",
		Expected:    "alyja/memory/memory-used",
		Packet: Packet{
			Hostname:     "alyja",
			Plugin:       "memory",
			Type:         "memory",
			TypeInstance: "used",
		},
	},
	{
		Description: "wanda/disk-hdc1/disk_octets",
		Expected:    "wanda/disk-hdc1/disk_octets",
		Packet: Packet{
			Hostname:       "wanda",
			Plugin:         "disk",
			PluginInstance: "hdc1",
			Type:           "disk_octets",
		},
	},
	{
		Description: "leeloo/load/load",
		Expected:    "leeloo/load/load",
		Packet: Packet{
			Hostname: "leeloo",
			Plugin:   "load",
			Type:     "load",
		},
	},
}

func TestFormat(t *testing.T) {
	assert := assert.New(t)

	for _, test := range formatTests {
		got := test.Packet.FormatName()
		assert.Equal(test.Expected, got, test.Description)
	}
}

func TestTypesDB(t *testing.T) {
	typedb, err := TypesDB(TypesDBData)
	if err != nil {
		t.Error(err)
	}
	dataSources, ok := typedb["compression"]
	if !ok {
		t.Error(`"compression" not found`)
	}

	found := len(dataSources)
	expected := 2
	if found != expected {
		t.Errorf("found %d datasources; expected %d", found, expected)
		return
	}

	ds := dataSources[0]

	if ds.Name != "uncompressed" {
		t.Errorf("ds.Name = %s; expected \"uncompressed\"", ds.Name)
	}

	if ds.Type != TypeDerive {
		dsType, ok := ValueTypeValues[ds.Type]
		if !ok {
			t.Error("unrecognized type")
		} else {
			t.Errorf("ds.Type = %s; expected TypeDerive", dsType)
		}
	}

	if ds.Min != "0" {
		t.Errorf("ds.Min = %s; expected 0", ds.Min)
	}

	if ds.Max != "U" {
		t.Errorf("ds.Max = %s; expected 0", ds.Max)
	}

	ds = dataSources[1]

	if ds.Name != "compressed" {
		t.Errorf("ds.Name = %s; expected \"uncompressed\"", ds.Name)
	}

	if ds.Type != TypeDerive {
		dsType, ok := ValueTypeValues[ds.Type]
		if !ok {
			t.Error("unrecognized type")
		} else {
			t.Errorf("ds.Type = %s; expected TypeDerive", dsType)
		}
	}

	if ds.Min != "0" {
		t.Errorf("ds.Min = %s; expected 0", ds.Min)
	}

	if ds.Max != "U" {
		t.Errorf("ds.Max = %s; expected 0", ds.Max)
	}
}

// Taken from /usr/share/collectd/types.db on a Ubuntu system
var TypesDBData = []byte(`absolute		value:ABSOLUTE:0:U
apache_bytes		value:DERIVE:0:U
apache_connections	value:GAUGE:0:65535
apache_idle_workers	value:GAUGE:0:65535
apache_requests		value:DERIVE:0:U
apache_scoreboard	value:GAUGE:0:65535
ath_nodes		value:GAUGE:0:65535
ath_stat		value:DERIVE:0:U
backends		value:GAUGE:0:65535
bitrate			value:GAUGE:0:4294967295
bytes			value:GAUGE:0:U
cache_eviction		value:DERIVE:0:U
cache_operation		value:DERIVE:0:U
cache_ratio		value:GAUGE:0:100
cache_result		value:DERIVE:0:U
cache_size		value:GAUGE:0:4294967295
charge			value:GAUGE:0:U
compression_ratio	value:GAUGE:0:2
compression		uncompressed:DERIVE:0:U, compressed:DERIVE:0:U
connections		value:DERIVE:0:U
conntrack		value:GAUGE:0:4294967295
contextswitch		value:DERIVE:0:U
counter			value:COUNTER:U:U
cpufreq			value:GAUGE:0:U
cpu			value:DERIVE:0:U
current_connections	value:GAUGE:0:U
current_sessions	value:GAUGE:0:U
current			value:GAUGE:U:U
delay			value:GAUGE:-1000000:1000000
derive			value:DERIVE:0:U
df_complex		value:GAUGE:0:U
df_inodes		value:GAUGE:0:U
df			used:GAUGE:0:1125899906842623, free:GAUGE:0:1125899906842623
disk_latency		read:GAUGE:0:U, write:GAUGE:0:U
disk_merged		read:DERIVE:0:U, write:DERIVE:0:U
disk_octets		read:DERIVE:0:U, write:DERIVE:0:U
disk_ops_complex	value:DERIVE:0:U
disk_ops		read:DERIVE:0:U, write:DERIVE:0:U
disk_time		read:DERIVE:0:U, write:DERIVE:0:U
dns_answer		value:DERIVE:0:U
dns_notify		value:DERIVE:0:U
dns_octets		queries:DERIVE:0:U, responses:DERIVE:0:U
dns_opcode		value:DERIVE:0:U
dns_qtype_cached	value:GAUGE:0:4294967295
dns_qtype		value:DERIVE:0:U
dns_query		value:DERIVE:0:U
dns_question		value:DERIVE:0:U
dns_rcode		value:DERIVE:0:U
dns_reject		value:DERIVE:0:U
dns_request		value:DERIVE:0:U
dns_resolver		value:DERIVE:0:U
dns_response		value:DERIVE:0:U
dns_transfer		value:DERIVE:0:U
dns_update		value:DERIVE:0:U
dns_zops		value:DERIVE:0:U
duration		seconds:GAUGE:0:U
email_check		value:GAUGE:0:U
email_count		value:GAUGE:0:U
email_size		value:GAUGE:0:U
entropy			value:GAUGE:0:4294967295
fanspeed		value:GAUGE:0:U
file_size		value:GAUGE:0:U
files			value:GAUGE:0:U
fork_rate		value:DERIVE:0:U
frequency_offset	value:GAUGE:-1000000:1000000
frequency		value:GAUGE:0:U
fscache_stat		value:DERIVE:0:U
gauge			value:GAUGE:U:U
hash_collisions		value:DERIVE:0:U
http_request_methods	value:DERIVE:0:U
http_requests		value:DERIVE:0:U
http_response_codes	value:DERIVE:0:U
humidity		value:GAUGE:0:100
if_collisions		value:DERIVE:0:U
if_dropped		rx:DERIVE:0:U, tx:DERIVE:0:U
if_errors		rx:DERIVE:0:U, tx:DERIVE:0:U
if_multicast		value:DERIVE:0:U
if_octets		rx:DERIVE:0:U, tx:DERIVE:0:U
if_packets		rx:DERIVE:0:U, tx:DERIVE:0:U
if_rx_errors		value:DERIVE:0:U
if_rx_octets		value:DERIVE:0:U
if_tx_errors		value:DERIVE:0:U
if_tx_octets		value:DERIVE:0:U
invocations		value:DERIVE:0:U
io_octets		rx:DERIVE:0:U, tx:DERIVE:0:U
io_packets		rx:DERIVE:0:U, tx:DERIVE:0:U
ipt_bytes		value:DERIVE:0:U
ipt_packets		value:DERIVE:0:U
irq			value:DERIVE:0:U
latency			value:GAUGE:0:U
links			value:GAUGE:0:U
load			shortterm:GAUGE:0:5000, midterm:GAUGE:0:5000, longterm:GAUGE:0:5000
md_disks		value:GAUGE:0:U
memcached_command	value:DERIVE:0:U
memcached_connections	value:GAUGE:0:U
memcached_items		value:GAUGE:0:U
memcached_octets	rx:DERIVE:0:U, tx:DERIVE:0:U
memcached_ops		value:DERIVE:0:U
memory			value:GAUGE:0:281474976710656
multimeter		value:GAUGE:U:U
mutex_operations	value:DERIVE:0:U
mysql_commands		value:DERIVE:0:U
mysql_handler		value:DERIVE:0:U
mysql_locks		value:DERIVE:0:U
mysql_log_position	value:DERIVE:0:U
mysql_octets		rx:DERIVE:0:U, tx:DERIVE:0:U
nfs_procedure		value:DERIVE:0:U
nginx_connections	value:GAUGE:0:U
nginx_requests		value:DERIVE:0:U
node_octets		rx:DERIVE:0:U, tx:DERIVE:0:U
node_rssi		value:GAUGE:0:255
node_stat		value:DERIVE:0:U
node_tx_rate		value:GAUGE:0:127
objects			value:GAUGE:0:U
operations		value:DERIVE:0:U
percent			value:GAUGE:0:100.1
percent_bytes		value:GAUGE:0:100.1
percent_inodes		value:GAUGE:0:100.1
pf_counters		value:DERIVE:0:U
pf_limits		value:DERIVE:0:U
pf_source		value:DERIVE:0:U
pf_states		value:GAUGE:0:U
pf_state		value:DERIVE:0:U
pg_blks			value:DERIVE:0:U
pg_db_size		value:GAUGE:0:U
pg_n_tup_c		value:DERIVE:0:U
pg_n_tup_g		value:GAUGE:0:U
pg_numbackends		value:GAUGE:0:U
pg_scan			value:DERIVE:0:U
pg_xact			value:DERIVE:0:U
ping_droprate		value:GAUGE:0:100
ping_stddev		value:GAUGE:0:65535
ping			value:GAUGE:0:65535
players			value:GAUGE:0:1000000
power			value:GAUGE:0:U
protocol_counter	value:DERIVE:0:U
ps_code			value:GAUGE:0:9223372036854775807
ps_count		processes:GAUGE:0:1000000, threads:GAUGE:0:1000000
ps_cputime		user:DERIVE:0:U, syst:DERIVE:0:U
ps_data			value:GAUGE:0:9223372036854775807
ps_disk_octets		read:DERIVE:0:U, write:DERIVE:0:U
ps_disk_ops		read:DERIVE:0:U, write:DERIVE:0:U
ps_pagefaults		minflt:DERIVE:0:U, majflt:DERIVE:0:U
ps_rss			value:GAUGE:0:9223372036854775807
ps_stacksize		value:GAUGE:0:9223372036854775807
ps_state		value:GAUGE:0:65535
ps_vm			value:GAUGE:0:9223372036854775807
queue_length		value:GAUGE:0:U
records			value:GAUGE:0:U
requests		value:GAUGE:0:U
response_time		value:GAUGE:0:U
response_code		value:GAUGE:0:U
route_etx		value:GAUGE:0:U
route_metric		value:GAUGE:0:U
routes			value:GAUGE:0:U
serial_octets		rx:DERIVE:0:U, tx:DERIVE:0:U
signal_noise		value:GAUGE:U:0
signal_power		value:GAUGE:U:0
signal_quality		value:GAUGE:0:U
snr			value:GAUGE:0:U
spam_check		value:GAUGE:0:U
spam_score		value:GAUGE:U:U
spl			value:GAUGE:U:U
swap_io			value:DERIVE:0:U
swap			value:GAUGE:0:1099511627776
tcp_connections		value:GAUGE:0:4294967295
temperature		value:GAUGE:U:U
threads			value:GAUGE:0:U
time_dispersion		value:GAUGE:-1000000:1000000
timeleft		value:GAUGE:0:U
time_offset		value:GAUGE:-1000000:1000000
total_bytes		value:DERIVE:0:U
total_connections	value:DERIVE:0:U
total_objects		value:DERIVE:0:U
total_operations	value:DERIVE:0:U
total_requests		value:DERIVE:0:U
total_sessions		value:DERIVE:0:U
total_threads		value:DERIVE:0:U
total_time_in_ms	value:DERIVE:0:U
total_values		value:DERIVE:0:U
uptime			value:GAUGE:0:4294967295
users			value:GAUGE:0:65535
vcl			value:GAUGE:0:65535
vcpu			value:GAUGE:0:U
virt_cpu_total		value:DERIVE:0:U
virt_vcpu		value:DERIVE:0:U
vmpage_action		value:DERIVE:0:U
vmpage_faults		minflt:DERIVE:0:U, majflt:DERIVE:0:U
vmpage_io		in:DERIVE:0:U, out:DERIVE:0:U
vmpage_number		value:GAUGE:0:4294967295
volatile_changes	value:GAUGE:0:U
voltage_threshold	value:GAUGE:U:U, threshold:GAUGE:U:U
voltage			value:GAUGE:U:U
vs_memory		value:GAUGE:0:9223372036854775807
vs_processes		value:GAUGE:0:65535
vs_threads		value:GAUGE:0:65535

#
# Legacy types
# (required for the v5 upgrade target)
#
arc_counts		demand_data:COUNTER:0:U, demand_metadata:COUNTER:0:U, prefetch_data:COUNTER:0:U, prefetch_metadata:COUNTER:0:U
arc_l2_bytes		read:COUNTER:0:U, write:COUNTER:0:U
arc_l2_size		value:GAUGE:0:U
arc_ratio		value:GAUGE:0:U
arc_size		current:GAUGE:0:U, target:GAUGE:0:U, minlimit:GAUGE:0:U, maxlimit:GAUGE:0:U
mysql_qcache		hits:COUNTER:0:U, inserts:COUNTER:0:U, not_cached:COUNTER:0:U, lowmem_prunes:COUNTER:0:U, queries_in_cache:GAUGE:0:U
mysql_threads		running:GAUGE:0:U, connected:GAUGE:0:U, cached:GAUGE:0:U, created:COUNTER:0:U
`)
