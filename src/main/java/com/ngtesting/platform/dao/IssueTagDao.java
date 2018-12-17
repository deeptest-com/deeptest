package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuAttachment;
import com.ngtesting.platform.model.IsuTag;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueTagDao {
    List<IsuTag> search(@Param("issueId") Integer issueId,
                        @Param("orgId") Integer orgId,
                        @Param("keywords") String keywords,
                        @Param("exceptIds")  List<Integer> exceptIds);

    List<IsuAttachment> query(@Param("tagId") Integer tagId);
    IsuAttachment get(@Param("id") Integer id);

    void save(@Param("list") List<IsuTag> list);

    void removeRelations(@Param("issueId") Integer issueId);
    void saveRelations(@Param("issueId") Integer issueId,
                       @Param("tags") List<IsuTag> tags);

    void updateTagField(@Param("issueId") Integer issueId,
                        @Param("tags") String tags);
}
