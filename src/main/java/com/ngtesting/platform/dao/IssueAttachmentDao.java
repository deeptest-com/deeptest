package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuAttachment;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueAttachmentDao {
    List<IsuAttachment> query(@Param("issueId") Integer issueId);
    IsuAttachment get(@Param("id") Integer id);

    void save(IsuAttachment attach);

    void delete(@Param("id") Integer id);
}
