package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuLinkReason;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueLinkDao {

    void link(@Param("srcIssueId") Integer srcIssueId,
              @Param("dictIssueId") Integer dictIssueId,
              @Param("reasonId") Integer reasonId,
              @Param("reasonName") String reasonName);

    List<IsuLinkReason> listLinkReason();
}
