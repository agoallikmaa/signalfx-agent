// Code generated by monitor-code-gen. DO NOT EDIT.

package mysql

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/mysql"

var groupSet = map[string]bool{}

const (
	cacheResultCacheSize             = "cache_result.cache_size"
	cacheResultQcacheHits            = "cache_result.qcache-hits"
	cacheResultQcacheInserts         = "cache_result.qcache-inserts"
	cacheResultQcacheNotCached       = "cache_result.qcache-not_cached"
	cacheResultQcachePrunes          = "cache_result.qcache-prunes"
	cacheSizeQcache                  = "cache_size.qcache"
	mysqlCommandsAdminCommands       = "mysql_commands.admin_commands"
	mysqlCommandsAlterDb             = "mysql_commands.alter_db"
	mysqlCommandsAlterDbUpgrade      = "mysql_commands.alter_db_upgrade"
	mysqlCommandsAlterEvent          = "mysql_commands.alter_event"
	mysqlCommandsAlterFunction       = "mysql_commands.alter_function"
	mysqlCommandsAlterProcedure      = "mysql_commands.alter_procedure"
	mysqlCommandsAlterServer         = "mysql_commands.alter_server"
	mysqlCommandsAlterTable          = "mysql_commands.alter_table"
	mysqlCommandsAlterTablespace     = "mysql_commands.alter_tablespace"
	mysqlCommandsAlterUser           = "mysql_commands.alter_user"
	mysqlCommandsAnalyze             = "mysql_commands.analyze"
	mysqlCommandsAssignToKeycache    = "mysql_commands.assign_to_keycache"
	mysqlCommandsBegin               = "mysql_commands.begin"
	mysqlCommandsBinlog              = "mysql_commands.binlog"
	mysqlCommandsCallProcedure       = "mysql_commands.call_procedure"
	mysqlCommandsChangeDb            = "mysql_commands.change_db"
	mysqlCommandsChangeMaster        = "mysql_commands.change_master"
	mysqlCommandsCheck               = "mysql_commands.check"
	mysqlCommandsChecksum            = "mysql_commands.checksum"
	mysqlCommandsCommit              = "mysql_commands.commit"
	mysqlCommandsCreateDb            = "mysql_commands.create_db"
	mysqlCommandsCreateEvent         = "mysql_commands.create_event"
	mysqlCommandsCreateFunction      = "mysql_commands.create_function"
	mysqlCommandsCreateIndex         = "mysql_commands.create_index"
	mysqlCommandsCreateProcedure     = "mysql_commands.create_procedure"
	mysqlCommandsCreateServer        = "mysql_commands.create_server"
	mysqlCommandsCreateTable         = "mysql_commands.create_table"
	mysqlCommandsCreateTrigger       = "mysql_commands.create_trigger"
	mysqlCommandsCreateUdf           = "mysql_commands.create_udf"
	mysqlCommandsCreateUser          = "mysql_commands.create_user"
	mysqlCommandsCreateView          = "mysql_commands.create_view"
	mysqlCommandsDeallocSQL          = "mysql_commands.dealloc_sql"
	mysqlCommandsDelete              = "mysql_commands.delete"
	mysqlCommandsDeleteMulti         = "mysql_commands.delete_multi"
	mysqlCommandsDo                  = "mysql_commands.do"
	mysqlCommandsDropDb              = "mysql_commands.drop_db"
	mysqlCommandsDropEvent           = "mysql_commands.drop_event"
	mysqlCommandsDropFunction        = "mysql_commands.drop_function"
	mysqlCommandsDropIndex           = "mysql_commands.drop_index"
	mysqlCommandsDropProcedure       = "mysql_commands.drop_procedure"
	mysqlCommandsDropServer          = "mysql_commands.drop_server"
	mysqlCommandsDropTable           = "mysql_commands.drop_table"
	mysqlCommandsDropTrigger         = "mysql_commands.drop_trigger"
	mysqlCommandsDropUser            = "mysql_commands.drop_user"
	mysqlCommandsDropView            = "mysql_commands.drop_view"
	mysqlCommandsEmptyQuery          = "mysql_commands.empty_query"
	mysqlCommandsExecuteSQL          = "mysql_commands.execute_sql"
	mysqlCommandsFlush               = "mysql_commands.flush"
	mysqlCommandsGetDiagnostics      = "mysql_commands.get_diagnostics"
	mysqlCommandsGrant               = "mysql_commands.grant"
	mysqlCommandsHaClose             = "mysql_commands.ha_close"
	mysqlCommandsHaOpen              = "mysql_commands.ha_open"
	mysqlCommandsHaRead              = "mysql_commands.ha_read"
	mysqlCommandsHelp                = "mysql_commands.help"
	mysqlCommandsInsert              = "mysql_commands.insert"
	mysqlCommandsInsertSelect        = "mysql_commands.insert_select"
	mysqlCommandsInstallPlugin       = "mysql_commands.install_plugin"
	mysqlCommandsKill                = "mysql_commands.kill"
	mysqlCommandsLoad                = "mysql_commands.load"
	mysqlCommandsLockTables          = "mysql_commands.lock_tables"
	mysqlCommandsOptimize            = "mysql_commands.optimize"
	mysqlCommandsPreloadKeys         = "mysql_commands.preload_keys"
	mysqlCommandsPrepareSQL          = "mysql_commands.prepare_sql"
	mysqlCommandsPurge               = "mysql_commands.purge"
	mysqlCommandsPurgeBeforeDate     = "mysql_commands.purge_before_date"
	mysqlCommandsReleaseSavepoint    = "mysql_commands.release_savepoint"
	mysqlCommandsRenameTable         = "mysql_commands.rename_table"
	mysqlCommandsRenameUser          = "mysql_commands.rename_user"
	mysqlCommandsRepair              = "mysql_commands.repair"
	mysqlCommandsReplace             = "mysql_commands.replace"
	mysqlCommandsReplaceSelect       = "mysql_commands.replace_select"
	mysqlCommandsReset               = "mysql_commands.reset"
	mysqlCommandsResignal            = "mysql_commands.resignal"
	mysqlCommandsRevoke              = "mysql_commands.revoke"
	mysqlCommandsRevokeAll           = "mysql_commands.revoke_all"
	mysqlCommandsRollback            = "mysql_commands.rollback"
	mysqlCommandsRollbackToSavepoint = "mysql_commands.rollback_to_savepoint"
	mysqlCommandsSavepoint           = "mysql_commands.savepoint"
	mysqlCommandsSelect              = "mysql_commands.select"
	mysqlCommandsSetOption           = "mysql_commands.set_option"
	mysqlCommandsShowBinlogEvents    = "mysql_commands.show_binlog_events"
	mysqlCommandsShowBinlogs         = "mysql_commands.show_binlogs"
	mysqlCommandsShowCharsets        = "mysql_commands.show_charsets"
	mysqlCommandsShowCollations      = "mysql_commands.show_collations"
	mysqlCommandsShowCreateDb        = "mysql_commands.show_create_db"
	mysqlCommandsShowCreateEvent     = "mysql_commands.show_create_event"
	mysqlCommandsShowCreateFunc      = "mysql_commands.show_create_func"
	mysqlCommandsShowCreateProc      = "mysql_commands.show_create_proc"
	mysqlCommandsShowCreateTable     = "mysql_commands.show_create_table"
	mysqlCommandsShowCreateTrigger   = "mysql_commands.show_create_trigger"
	mysqlCommandsShowDatabases       = "mysql_commands.show_databases"
	mysqlCommandsShowEngineLogs      = "mysql_commands.show_engine_logs"
	mysqlCommandsShowEngineMutex     = "mysql_commands.show_engine_mutex"
	mysqlCommandsShowEngineStatus    = "mysql_commands.show_engine_status"
	mysqlCommandsShowErrors          = "mysql_commands.show_errors"
	mysqlCommandsShowEvents          = "mysql_commands.show_events"
	mysqlCommandsShowFields          = "mysql_commands.show_fields"
	mysqlCommandsShowFunctionCode    = "mysql_commands.show_function_code"
	mysqlCommandsShowFunctionStatus  = "mysql_commands.show_function_status"
	mysqlCommandsShowGrants          = "mysql_commands.show_grants"
	mysqlCommandsShowKeys            = "mysql_commands.show_keys"
	mysqlCommandsShowMasterStatus    = "mysql_commands.show_master_status"
	mysqlCommandsShowOpenTables      = "mysql_commands.show_open_tables"
	mysqlCommandsShowPlugins         = "mysql_commands.show_plugins"
	mysqlCommandsShowPrivileges      = "mysql_commands.show_privileges"
	mysqlCommandsShowProcedureCode   = "mysql_commands.show_procedure_code"
	mysqlCommandsShowProcedureStatus = "mysql_commands.show_procedure_status"
	mysqlCommandsShowProcesslist     = "mysql_commands.show_processlist"
	mysqlCommandsShowProfile         = "mysql_commands.show_profile"
	mysqlCommandsShowProfiles        = "mysql_commands.show_profiles"
	mysqlCommandsShowRelaylogEvents  = "mysql_commands.show_relaylog_events"
	mysqlCommandsShowSlaveHosts      = "mysql_commands.show_slave_hosts"
	mysqlCommandsShowSlaveStatus     = "mysql_commands.show_slave_status"
	mysqlCommandsShowStatus          = "mysql_commands.show_status"
	mysqlCommandsShowStorageEngines  = "mysql_commands.show_storage_engines"
	mysqlCommandsShowTableStatus     = "mysql_commands.show_table_status"
	mysqlCommandsShowTables          = "mysql_commands.show_tables"
	mysqlCommandsShowTriggers        = "mysql_commands.show_triggers"
	mysqlCommandsShowVariables       = "mysql_commands.show_variables"
	mysqlCommandsShowWarnings        = "mysql_commands.show_warnings"
	mysqlCommandsSignal              = "mysql_commands.signal"
	mysqlCommandsSlaveStart          = "mysql_commands.slave_start"
	mysqlCommandsSlaveStop           = "mysql_commands.slave_stop"
	mysqlCommandsTruncate            = "mysql_commands.truncate"
	mysqlCommandsUninstallPlugin     = "mysql_commands.uninstall_plugin"
	mysqlCommandsUnlockTables        = "mysql_commands.unlock_tables"
	mysqlCommandsUpdate              = "mysql_commands.update"
	mysqlCommandsUpdateMulti         = "mysql_commands.update_multi"
	mysqlCommandsXaCommit            = "mysql_commands.xa_commit"
	mysqlCommandsXaEnd               = "mysql_commands.xa_end"
	mysqlCommandsXaPrepare           = "mysql_commands.xa_prepare"
	mysqlCommandsXaRecover           = "mysql_commands.xa_recover"
	mysqlCommandsXaRollback          = "mysql_commands.xa_rollback"
	mysqlCommandsXaStart             = "mysql_commands.xa_start"
	mysqlHandlerCommit               = "mysql_handler.commit"
	mysqlHandlerDelete               = "mysql_handler.delete"
	mysqlHandlerExternalLock         = "mysql_handler.external_lock"
	mysqlHandlerPrepare              = "mysql_handler.prepare"
	mysqlHandlerReadFirst            = "mysql_handler.read_first"
	mysqlHandlerReadKey              = "mysql_handler.read_key"
	mysqlHandlerReadNext             = "mysql_handler.read_next"
	mysqlHandlerReadPrev             = "mysql_handler.read_prev"
	mysqlHandlerReadRnd              = "mysql_handler.read_rnd"
	mysqlHandlerReadRndNext          = "mysql_handler.read_rnd_next"
	mysqlHandlerRollback             = "mysql_handler.rollback"
	mysqlHandlerSavepoint            = "mysql_handler.savepoint"
	mysqlHandlerSavepointRollback    = "mysql_handler.savepoint_rollback"
	mysqlHandlerUpdate               = "mysql_handler.update"
	mysqlHandlerWrite                = "mysql_handler.write"
	mysqlLocksImmediate              = "mysql_locks.immediate"
	mysqlLocksWaited                 = "mysql_locks.waited"
	mysqlOctetsRx                    = "mysql_octets.rx"
	mysqlOctetsTx                    = "mysql_octets.tx"
	mysqlSelectFullJoin              = "mysql_select.full_join"
	mysqlSelectFullRangeJoin         = "mysql_select.full_range_join"
	mysqlSelectRange                 = "mysql_select.range"
	mysqlSelectRangeCheck            = "mysql_select.range_check"
	mysqlSelectScan                  = "mysql_select.scan"
	mysqlSlowQueries                 = "mysql_slow_queries"
	mysqlSortRange                   = "mysql_sort.range"
	mysqlSortScan                    = "mysql_sort.scan"
	mysqlSortMergePasses             = "mysql_sort_merge_passes"
	mysqlSortRows                    = "mysql_sort_rows"
	threadsCached                    = "threads.cached"
	threadsConnected                 = "threads.connected"
	threadsRunning                   = "threads.running"
	totalThreadsCreated              = "total_threads.created"
)

var metricSet = map[string]monitors.MetricInfo{
	cacheResultCacheSize:             {Type: datapoint.Gauge},
	cacheResultQcacheHits:            {Type: datapoint.Counter},
	cacheResultQcacheInserts:         {Type: datapoint.Counter},
	cacheResultQcacheNotCached:       {Type: datapoint.Counter},
	cacheResultQcachePrunes:          {Type: datapoint.Counter},
	cacheSizeQcache:                  {Type: datapoint.Gauge},
	mysqlCommandsAdminCommands:       {Type: datapoint.Counter},
	mysqlCommandsAlterDb:             {Type: datapoint.Counter},
	mysqlCommandsAlterDbUpgrade:      {Type: datapoint.Counter},
	mysqlCommandsAlterEvent:          {Type: datapoint.Counter},
	mysqlCommandsAlterFunction:       {Type: datapoint.Counter},
	mysqlCommandsAlterProcedure:      {Type: datapoint.Counter},
	mysqlCommandsAlterServer:         {Type: datapoint.Counter},
	mysqlCommandsAlterTable:          {Type: datapoint.Counter},
	mysqlCommandsAlterTablespace:     {Type: datapoint.Counter},
	mysqlCommandsAlterUser:           {Type: datapoint.Counter},
	mysqlCommandsAnalyze:             {Type: datapoint.Counter},
	mysqlCommandsAssignToKeycache:    {Type: datapoint.Counter},
	mysqlCommandsBegin:               {Type: datapoint.Counter},
	mysqlCommandsBinlog:              {Type: datapoint.Counter},
	mysqlCommandsCallProcedure:       {Type: datapoint.Counter},
	mysqlCommandsChangeDb:            {Type: datapoint.Counter},
	mysqlCommandsChangeMaster:        {Type: datapoint.Counter},
	mysqlCommandsCheck:               {Type: datapoint.Counter},
	mysqlCommandsChecksum:            {Type: datapoint.Counter},
	mysqlCommandsCommit:              {Type: datapoint.Counter},
	mysqlCommandsCreateDb:            {Type: datapoint.Counter},
	mysqlCommandsCreateEvent:         {Type: datapoint.Counter},
	mysqlCommandsCreateFunction:      {Type: datapoint.Counter},
	mysqlCommandsCreateIndex:         {Type: datapoint.Counter},
	mysqlCommandsCreateProcedure:     {Type: datapoint.Counter},
	mysqlCommandsCreateServer:        {Type: datapoint.Counter},
	mysqlCommandsCreateTable:         {Type: datapoint.Counter},
	mysqlCommandsCreateTrigger:       {Type: datapoint.Counter},
	mysqlCommandsCreateUdf:           {Type: datapoint.Counter},
	mysqlCommandsCreateUser:          {Type: datapoint.Counter},
	mysqlCommandsCreateView:          {Type: datapoint.Counter},
	mysqlCommandsDeallocSQL:          {Type: datapoint.Counter},
	mysqlCommandsDelete:              {Type: datapoint.Counter},
	mysqlCommandsDeleteMulti:         {Type: datapoint.Counter},
	mysqlCommandsDo:                  {Type: datapoint.Counter},
	mysqlCommandsDropDb:              {Type: datapoint.Counter},
	mysqlCommandsDropEvent:           {Type: datapoint.Counter},
	mysqlCommandsDropFunction:        {Type: datapoint.Counter},
	mysqlCommandsDropIndex:           {Type: datapoint.Counter},
	mysqlCommandsDropProcedure:       {Type: datapoint.Counter},
	mysqlCommandsDropServer:          {Type: datapoint.Counter},
	mysqlCommandsDropTable:           {Type: datapoint.Counter},
	mysqlCommandsDropTrigger:         {Type: datapoint.Counter},
	mysqlCommandsDropUser:            {Type: datapoint.Counter},
	mysqlCommandsDropView:            {Type: datapoint.Counter},
	mysqlCommandsEmptyQuery:          {Type: datapoint.Counter},
	mysqlCommandsExecuteSQL:          {Type: datapoint.Counter},
	mysqlCommandsFlush:               {Type: datapoint.Counter},
	mysqlCommandsGetDiagnostics:      {Type: datapoint.Counter},
	mysqlCommandsGrant:               {Type: datapoint.Counter},
	mysqlCommandsHaClose:             {Type: datapoint.Counter},
	mysqlCommandsHaOpen:              {Type: datapoint.Counter},
	mysqlCommandsHaRead:              {Type: datapoint.Counter},
	mysqlCommandsHelp:                {Type: datapoint.Counter},
	mysqlCommandsInsert:              {Type: datapoint.Counter},
	mysqlCommandsInsertSelect:        {Type: datapoint.Counter},
	mysqlCommandsInstallPlugin:       {Type: datapoint.Counter},
	mysqlCommandsKill:                {Type: datapoint.Counter},
	mysqlCommandsLoad:                {Type: datapoint.Counter},
	mysqlCommandsLockTables:          {Type: datapoint.Counter},
	mysqlCommandsOptimize:            {Type: datapoint.Counter},
	mysqlCommandsPreloadKeys:         {Type: datapoint.Counter},
	mysqlCommandsPrepareSQL:          {Type: datapoint.Counter},
	mysqlCommandsPurge:               {Type: datapoint.Counter},
	mysqlCommandsPurgeBeforeDate:     {Type: datapoint.Counter},
	mysqlCommandsReleaseSavepoint:    {Type: datapoint.Counter},
	mysqlCommandsRenameTable:         {Type: datapoint.Counter},
	mysqlCommandsRenameUser:          {Type: datapoint.Counter},
	mysqlCommandsRepair:              {Type: datapoint.Counter},
	mysqlCommandsReplace:             {Type: datapoint.Counter},
	mysqlCommandsReplaceSelect:       {Type: datapoint.Counter},
	mysqlCommandsReset:               {Type: datapoint.Counter},
	mysqlCommandsResignal:            {Type: datapoint.Counter},
	mysqlCommandsRevoke:              {Type: datapoint.Counter},
	mysqlCommandsRevokeAll:           {Type: datapoint.Counter},
	mysqlCommandsRollback:            {Type: datapoint.Counter},
	mysqlCommandsRollbackToSavepoint: {Type: datapoint.Counter},
	mysqlCommandsSavepoint:           {Type: datapoint.Counter},
	mysqlCommandsSelect:              {Type: datapoint.Counter},
	mysqlCommandsSetOption:           {Type: datapoint.Counter},
	mysqlCommandsShowBinlogEvents:    {Type: datapoint.Counter},
	mysqlCommandsShowBinlogs:         {Type: datapoint.Counter},
	mysqlCommandsShowCharsets:        {Type: datapoint.Counter},
	mysqlCommandsShowCollations:      {Type: datapoint.Counter},
	mysqlCommandsShowCreateDb:        {Type: datapoint.Counter},
	mysqlCommandsShowCreateEvent:     {Type: datapoint.Counter},
	mysqlCommandsShowCreateFunc:      {Type: datapoint.Counter},
	mysqlCommandsShowCreateProc:      {Type: datapoint.Counter},
	mysqlCommandsShowCreateTable:     {Type: datapoint.Counter},
	mysqlCommandsShowCreateTrigger:   {Type: datapoint.Counter},
	mysqlCommandsShowDatabases:       {Type: datapoint.Counter},
	mysqlCommandsShowEngineLogs:      {Type: datapoint.Counter},
	mysqlCommandsShowEngineMutex:     {Type: datapoint.Counter},
	mysqlCommandsShowEngineStatus:    {Type: datapoint.Counter},
	mysqlCommandsShowErrors:          {Type: datapoint.Counter},
	mysqlCommandsShowEvents:          {Type: datapoint.Counter},
	mysqlCommandsShowFields:          {Type: datapoint.Counter},
	mysqlCommandsShowFunctionCode:    {Type: datapoint.Counter},
	mysqlCommandsShowFunctionStatus:  {Type: datapoint.Counter},
	mysqlCommandsShowGrants:          {Type: datapoint.Counter},
	mysqlCommandsShowKeys:            {Type: datapoint.Counter},
	mysqlCommandsShowMasterStatus:    {Type: datapoint.Counter},
	mysqlCommandsShowOpenTables:      {Type: datapoint.Counter},
	mysqlCommandsShowPlugins:         {Type: datapoint.Counter},
	mysqlCommandsShowPrivileges:      {Type: datapoint.Counter},
	mysqlCommandsShowProcedureCode:   {Type: datapoint.Counter},
	mysqlCommandsShowProcedureStatus: {Type: datapoint.Counter},
	mysqlCommandsShowProcesslist:     {Type: datapoint.Counter},
	mysqlCommandsShowProfile:         {Type: datapoint.Counter},
	mysqlCommandsShowProfiles:        {Type: datapoint.Counter},
	mysqlCommandsShowRelaylogEvents:  {Type: datapoint.Counter},
	mysqlCommandsShowSlaveHosts:      {Type: datapoint.Counter},
	mysqlCommandsShowSlaveStatus:     {Type: datapoint.Counter},
	mysqlCommandsShowStatus:          {Type: datapoint.Counter},
	mysqlCommandsShowStorageEngines:  {Type: datapoint.Counter},
	mysqlCommandsShowTableStatus:     {Type: datapoint.Counter},
	mysqlCommandsShowTables:          {Type: datapoint.Counter},
	mysqlCommandsShowTriggers:        {Type: datapoint.Counter},
	mysqlCommandsShowVariables:       {Type: datapoint.Counter},
	mysqlCommandsShowWarnings:        {Type: datapoint.Counter},
	mysqlCommandsSignal:              {Type: datapoint.Counter},
	mysqlCommandsSlaveStart:          {Type: datapoint.Counter},
	mysqlCommandsSlaveStop:           {Type: datapoint.Counter},
	mysqlCommandsTruncate:            {Type: datapoint.Counter},
	mysqlCommandsUninstallPlugin:     {Type: datapoint.Counter},
	mysqlCommandsUnlockTables:        {Type: datapoint.Counter},
	mysqlCommandsUpdate:              {Type: datapoint.Counter},
	mysqlCommandsUpdateMulti:         {Type: datapoint.Counter},
	mysqlCommandsXaCommit:            {Type: datapoint.Counter},
	mysqlCommandsXaEnd:               {Type: datapoint.Counter},
	mysqlCommandsXaPrepare:           {Type: datapoint.Counter},
	mysqlCommandsXaRecover:           {Type: datapoint.Counter},
	mysqlCommandsXaRollback:          {Type: datapoint.Counter},
	mysqlCommandsXaStart:             {Type: datapoint.Counter},
	mysqlHandlerCommit:               {Type: datapoint.Counter},
	mysqlHandlerDelete:               {Type: datapoint.Counter},
	mysqlHandlerExternalLock:         {Type: datapoint.Counter},
	mysqlHandlerPrepare:              {Type: datapoint.Counter},
	mysqlHandlerReadFirst:            {Type: datapoint.Counter},
	mysqlHandlerReadKey:              {Type: datapoint.Counter},
	mysqlHandlerReadNext:             {Type: datapoint.Counter},
	mysqlHandlerReadPrev:             {Type: datapoint.Counter},
	mysqlHandlerReadRnd:              {Type: datapoint.Counter},
	mysqlHandlerReadRndNext:          {Type: datapoint.Counter},
	mysqlHandlerRollback:             {Type: datapoint.Counter},
	mysqlHandlerSavepoint:            {Type: datapoint.Counter},
	mysqlHandlerSavepointRollback:    {Type: datapoint.Counter},
	mysqlHandlerUpdate:               {Type: datapoint.Counter},
	mysqlHandlerWrite:                {Type: datapoint.Counter},
	mysqlLocksImmediate:              {Type: datapoint.Counter},
	mysqlLocksWaited:                 {Type: datapoint.Counter},
	mysqlOctetsRx:                    {Type: datapoint.Counter},
	mysqlOctetsTx:                    {Type: datapoint.Counter},
	mysqlSelectFullJoin:              {Type: datapoint.Counter},
	mysqlSelectFullRangeJoin:         {Type: datapoint.Counter},
	mysqlSelectRange:                 {Type: datapoint.Counter},
	mysqlSelectRangeCheck:            {Type: datapoint.Counter},
	mysqlSelectScan:                  {Type: datapoint.Counter},
	mysqlSlowQueries:                 {Type: datapoint.Counter},
	mysqlSortRange:                   {Type: datapoint.Counter},
	mysqlSortScan:                    {Type: datapoint.Counter},
	mysqlSortMergePasses:             {Type: datapoint.Counter},
	mysqlSortRows:                    {Type: datapoint.Counter},
	threadsCached:                    {Type: datapoint.Gauge},
	threadsConnected:                 {Type: datapoint.Gauge},
	threadsRunning:                   {Type: datapoint.Gauge},
	totalThreadsCreated:              {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	cacheResultQcacheHits:    true,
	cacheResultQcacheInserts: true,
	cacheSizeQcache:          true,
	mysqlCommandsDelete:      true,
	mysqlCommandsInsert:      true,
	mysqlCommandsSelect:      true,
	mysqlCommandsUpdate:      true,
	mysqlLocksImmediate:      true,
	mysqlLocksWaited:         true,
	mysqlOctetsRx:            true,
	mysqlOctetsTx:            true,
	threadsCached:            true,
	threadsConnected:         true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/mysql",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}