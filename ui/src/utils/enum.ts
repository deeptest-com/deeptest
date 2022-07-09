export enum ProcessorCategory {
    ProcessorThread = "processor_thread",
    ProcessorSimple = "processor_simple",
    ProcessorLogic = "processor_logic",
    ProcessorLoop = "processor_loop",
    ProcessorTimer = "processor_timer",
    ProcessorVariable  = "processor_variable",
    ProcessorAssertion = "processor_assertion",
    ProcessorExtractor = "processor_extractor",

    ProcessorCookie = "processor_cookie",
    ProcessorData = "processor_data",
}

export enum ProcessorThread {
    Thread = "processor_thread_default",
}
export enum ProcessorSimple {
    Simple = "processor_simple_default",
}
export enum ProcessorTimer {
    Time = "processor_time_default",
}

export enum ProcessorLogic {
    If = "processor_logic_if",
    Else = "processor_logic_else",
}

export enum ProcessorLoop {
    RepeatTime = "processor_loop_time",
    RepeatWhile = "processor_loop_while",
    RepeatIn = "processor_loop_in",
    RepeatRange = "processor_loop_range",
    RepeatBreak = "processor_loop_break",
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
    ZenData = "processor_data_zendata",
}