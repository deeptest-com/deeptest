package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseStep;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseStepDao {
    List<TstCaseStep> query(@Param("caseId") Integer caseId);
    TstCaseStep get(@Param("id") Integer id);
    void save(TstCaseStep vo);
    void update(TstCaseStep vo);
    void delete(@Param("id") Integer id);

    void moveOthersUp(@Param("caseId") Integer caseId, @Param("ordr") Integer ordr);
    void moveOthersDown(@Param("caseId") Integer caseId, @Param("id") Integer id, @Param("ordr") Integer ordr);

    TstCaseStep getPrev(@Param("caseId") Integer caseId, @Param("ordr") Integer ordr);
    TstCaseStep getNext(@Param("caseId") Integer caseId, @Param("ordr") Integer ordr);
    void setOrder(@Param("id") Integer id, @Param("ordr") Integer ordr);
}
