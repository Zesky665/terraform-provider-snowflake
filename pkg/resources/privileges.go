package resources

type Privilege string

func (p Privilege) String() string {
	return string(p)
}

const (
	privilegeAccountSupportCases       Privilege = "MANAGE ACCOUNT SUPPORT CASES"
	privilegeAddSearchOptimization     Privilege = "ADD SEARCH OPTIMIZATION"
	privilegeApply                     Privilege = "APPLY"
	privilegeApplyMaskingPolicy        Privilege = "APPLY MASKING POLICY"
	privilegeApplyPasswordPolicy       Privilege = "APPLY PASSWORD POLICY"
	privilegeApplyRowAccessPolicy      Privilege = "APPLY ROW ACCESS POLICY"
	privilegeApplySessionPolicy        Privilege = "APPLY SESSION POLICY"
	privilegeApplyTag                  Privilege = "APPLY TAG"
	privilegeAttachPolicy              Privilege = "ATTACH POLICY"
	privilegeCreateAccount             Privilege = "CREATE ACCOUNT"
	privilegeCreateDatabase            Privilege = "CREATE DATABASE"
	privilegeCreateDataExchangeListing Privilege = "CREATE DATA EXCHANGE LISTING"
	privilegeCreateExternalTable       Privilege = "CREATE EXTERNAL TABLE"
	privilegeCreateFileFormat          Privilege = "CREATE FILE FORMAT"
	privilegeCreateFunction            Privilege = "CREATE FUNCTION"
	privilegeCreateIntegration         Privilege = "CREATE INTEGRATION"
	privilegeCreateMaskingPolicy       Privilege = "CREATE MASKING POLICY"
	privilegeCreateMaterializedView    Privilege = "CREATE MATERIALIZED VIEW"
	privilegeCreateNetworkPolicy       Privilege = "CREATE NETWORK POLICY"
	privilegeCreatePipe                Privilege = "CREATE PIPE"
	privilegeCreateProcedure           Privilege = "CREATE PROCEDURE"
	privilegeCreateRole                Privilege = "CREATE ROLE"
	privilegeCreateRowAccessPolicy     Privilege = "CREATE ROW ACCESS POLICY"
	privilegeCreateSchema              Privilege = "CREATE SCHEMA"
	privilegeCreateSequence            Privilege = "CREATE SEQUENCE"
	privilegeCreateSessionPolicy       Privilege = "CREATE SESSION POLICY"
	privilegeCreateShare               Privilege = "CREATE SHARE"
	privilegeCreateStage               Privilege = "CREATE STAGE"
	privilegeCreateStream              Privilege = "CREATE STREAM"
	privilegeCreateTable               Privilege = "CREATE TABLE"
	privilegeCreateTag                 Privilege = "CREATE TAG"
	privilegeCreateTask                Privilege = "CREATE TASK"
	privilegeCreateTemporaryTable      Privilege = "CREATE TEMPORARY TABLE"
	privilegeCreateUser                Privilege = "CREATE USER"
	privilegeCreateView                Privilege = "CREATE VIEW"
	privilegeCreateWarehouse           Privilege = "CREATE WAREHOUSE"
	privilegeDelete                    Privilege = "DELETE"
	privilegeExecuteManagedTask        Privilege = "EXECUTE MANAGED TASK"
	privilegeExecuteTask               Privilege = "EXECUTE TASK"
	privilegeImportedPrivileges        Privilege = "IMPORTED PRIVILEGES"
	privilegeImportShare               Privilege = "IMPORT SHARE"
	privilegeInsert                    Privilege = "INSERT"
	privilegeManageGrants              Privilege = "MANAGE GRANTS"
	privilegeModify                    Privilege = "MODIFY"
	privilegeMonitor                   Privilege = "MONITOR"
	privilegeMonitorExecution          Privilege = "MONITOR EXECUTION"
	privilegeMonitorUsage              Privilege = "MONITOR USAGE"
	privilegeOperate                   Privilege = "OPERATE"
	privilegeOrganizationSupportCases  Privilege = "MANAGE ORGANIZATION SUPPORT CASES"
	privilegeOverrideShareRestrictions Privilege = "OVERRIDE SHARE RESTRICTIONS"
	privilegeOwnership                 Privilege = "OWNERSHIP"
	privilegeRead                      Privilege = "READ"
	privilegeRebuild                   Privilege = "REBUILD"
	privilegeReferences                Privilege = "REFERENCES"
	privilegeReferenceUsage            Privilege = "REFERENCE_USAGE"
	privilegeSelect                    Privilege = "SELECT"
	privilegeTruncate                  Privilege = "TRUNCATE"
	privilegeUpdate                    Privilege = "UPDATE"
	privilegeUsage                     Privilege = "USAGE"
	privilegeUserSupportCases          Privilege = "MANAGE USER SUPPORT CASES"
	privilegeWrite                     Privilege = "WRITE"
)

type PrivilegeSet map[Privilege]struct{}

func NewPrivilegeSet(privileges ...Privilege) PrivilegeSet {
	ps := PrivilegeSet{}
	for _, priv := range privileges {
		ps[priv] = struct{}{}
	}
	return ps
}

func (ps PrivilegeSet) ToList() []string {
	privs := []string{}
	for p := range ps {
		privs = append(privs, string(p))
	}
	return privs
}

func (ps PrivilegeSet) addString(s string) {
	ps[Privilege(s)] = struct{}{}
}

func (ps PrivilegeSet) hasString(s string) bool {
	_, ok := ps[Privilege(s)]
	return ok
}
