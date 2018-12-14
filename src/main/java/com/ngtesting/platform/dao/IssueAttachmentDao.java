package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuAttachment;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueAttachmentDao {
    List<IsuAttachment> query(@Param("caseId") Integer caseId);
    IsuAttachment get(@Param("id") Integer id);

    void save(IsuAttachment attach);

    void delete(@Param("id") Integer id);
}
