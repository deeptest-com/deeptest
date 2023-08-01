export enum UsedBy {
    InterfaceDebug = "interface_debug",
    ScenarioDebug = "scenario_debug",
    DiagnoseDebug = "diagnose_debug",
    CaseDebug = "case_debug",
}

export enum VarScope {
    ScopePublic = "public",
    ScopePrivate = "private",
}

export enum WsMsgCategory {
    InProgress = "in_progress",
    End = "end",
    Result = "result",
}

export enum ProcessorAction {
    ActionEdit = 'action_edit',
    ActionRemove = 'action_remove',
    ActionImportInterface = 'action_import_interface',
    ActionAddProcessor = 'action_add_processor',
    ActionInInterface = 'action_in_interface',
}

export enum ProcessorCategory {
    ProcessorRoot = "processor_root",
    ProcessorThread = "processor_thread",
    ProcessorGroup = "processor_group",
    ProcessorInterface = "processor_interface",
    ProcessorLogic = "processor_logic",
    ProcessorLoop = "processor_loop",
    ProcessorData = "processor_data",
    ProcessorVariable  = "processor_variable",
    ProcessorCookie = "processor_cookie",
    ProcessorExtractor = "processor_extractor",
    ProcessorTimer = "processor_timer",
    ProcessorPrint = "processor_print",
    ProcessorAssertion = "processor_assertion",
}
export enum ProcessorInterface {
    Interface = "processor_interface_default",
}
export enum ProcessorRoot {
    Root = "processor_root_default",
}
export enum ProcessorThread {
    Thread = "processor_thread_default",
}
export enum ProcessorGroup {
    Group = "processor_group_default",
}
export enum ProcessorTimer {
    Time = "processor_time_default",
}
export enum ProcessorPrint {
    Print = "processor_print_default",
}

export enum ProcessorLogic {
    If = "processor_logic_if",
    Else = "processor_logic_else",
}

export enum ProcessorLoop {
    Time = "processor_loop_time",
    Until = "processor_loop_until",
    In = "processor_loop_in",
    Range = "processor_loop_range",
    Break = "processor_loop_break",
}

export enum ProcessorVariable {
    // Get = "processor_variable_get",
    Set = "processor_variable_set",
    Clear = "processor_variable_clear",
}

export enum ProcessorAssertion {
    Assertion = "processor_assertion_default",
    // Equal      = "processor_assertion_equal",
    // NotEqual   = "processor_assertion_not_equal",
    // Contain    = "processor_assertion_contain",
    // NotContain = "processor_assertion_not_contain"
}

export enum ProcessorExtractor {
    Boundary = "processor_extractor_boundary",
    JsonQuery = "processor_extractor_jsonquery",
    HtmlQuery = "processor_extractor_htmlquery",
    XmlQuery = "processor_extractor_xmlquery",
}

export enum ProcessorCookie {
    Get = "processor_cookie_get",
    Set = "processor_cookie_set",
    Clear = "processor_cookie_clear",
}

export enum ProcessorData {
    Text = "processor_data_text",
    Excel = "processor_data_excel",
    // ZenData = "processor_data_zendata",
}

export enum RequestBodyType {
    'application/json'= 'application/json',
    'application/xml' = 'application/xml',
    'multipart/form-data' = 'multipart/form-data',
    'application/x-www-form-urlencoded' = 'application/x-www-form-urlencoded',
    'text/html' = 'text/html',
    'text/plain' = 'text/plain',
}

export const Methods = [
    "GET",
    "POST",
    "PUT",
    "PATCH",
    "DELETE",
    "HEAD",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "CUSTOM",
]

export enum ComparisonOperator {
    empty = 'empty',
    equal = 'equal',
    notEqual = 'notEqual',
    greaterThan = 'greaterThan',
    lessThan = 'lessThan',
    greaterThanOrEqual = 'greaterThanOrEqual',
    lessThanOrEqual = 'lessThanOrEqual',

    contain = 'contain',
    notContain = 'notContain',
}

export enum ConditionType {
    extractor = 'extractor',
    checkpoint = 'checkpoint',
    script = 'script',
}

export enum ExtractorSrc {
    header = 'header',
    body = 'body',
}
export enum ExtractorType {
    boundary = 'boundary',
    jsonquery = 'jsonquery',
    htmlquery = 'htmlquery',
    xmlquery = 'xmlquery',
    regx = 'regx',
    // fulltext = 'fulltext',
}
export enum CheckpointType {
    responseStatus = 'responseStatus',
    responseHeader = 'responseHeader',
    responseBody = 'responseBody',
    extractor = 'extractor',
    judgement = 'judgement'
}

export enum AuthorizationTypes {
    '' = 'None',
    'basicAuth' = 'Basic Auth',
    'bearerToken' = 'Bearer Token',
    'apiKey' = 'API Key',
    // 'oAuth2' = 'OAuth 2.0',
}

export enum OAuth2GrantTypes {
    'authorizationCode' = 'Authorization Code',
    'authorizationCodeWithPKCE' = 'Authorization Code (With PKCE)',
    'implicit' = 'Implicit',
    'passwordCredential' = 'Password Credential',
    'clientCredential' = 'Client Credential',
}

export enum OAuth2ClientAuthenticationWay {
    'sendAsBasicAuthHeader' = 'Send as Basic Auth header',
    'sendClientCredentialsInBody' = 'Send client credentials in body',
}

export enum ReportDetailType {
    ExecPlan = 'exec_plan',
    ExecScenario = 'exec_scenario',
    QueryDetail = 'query_detail'
}

export enum ProjectType {
    Full = 'full',
    Debug = 'debug'
}