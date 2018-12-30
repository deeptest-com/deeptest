package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseAttachment;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseAttachmentDao {
    List<TstCaseAttachment> query(@Param("caseId") Integer caseId);
    TstCaseAttachment get(@Param("id") Integer id);

    void save(TstCaseAttachment attach);

    void delete(@Param("id") Integer id,
                @Param("userId") Integer userId);
}
