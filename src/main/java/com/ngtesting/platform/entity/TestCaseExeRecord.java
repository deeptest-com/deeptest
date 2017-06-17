package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_exe_record")
public class TestCaseExeRecord extends TestCase {

    private String statusCode;
    private String statusName;

    @Column(name = "run_id")
    private Long runId;

    @Column(name = "case_id")
    private Long caseId;

}
