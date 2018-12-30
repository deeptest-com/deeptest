package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInTaskAttachment;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseInTaskAttachmentDao {
    List<TstCaseInTaskAttachment> query(@Param("caseInTaskId") Integer caseInTaskId);
    TstCaseInTaskAttachment get(@Param("id") Integer id);

    void save(TstCaseInTaskAttachment attach);

    void delete(@Param("id") Integer id,
                @Param("userId") Integer userId);
}
