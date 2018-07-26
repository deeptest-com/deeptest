package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseStep;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseStepDao {
    List<TstCaseStep> query(@Param("caseId") Integer caseId);
    void save(TstCaseStep step1);
}
