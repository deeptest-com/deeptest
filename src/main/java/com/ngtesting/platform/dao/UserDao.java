package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface UserDao {
    List<TstUser> query(@Param("orgId") Integer orgId,
                        @Param("keywords") String keywords,
                        @Param("disabled") Boolean disabled);

    List<TstUser> search(@Param("orgId") Integer orgId,
                         @Param("keywords") String keywords,
                         @Param("exceptIds") String exceptIds);

    TstUser get(@Param("id") Integer id);
    TstUser getByEmail(@Param("email") String email);
    TstUser getByEmailAndPassword(@Param("email") String email,
                                  @Param("password") String password);
    TstUser getByToken(@Param("token") String token);
    TstUser getByPhone(@Param("phone") String phone);

    String getSalt(@Param("id") Integer id);

    void save(TstUser record);
    void saveSettings(TstUser record);
    void update(TstUser record);

    void modifyProp(@Param("id") Integer id,
                    @Param("prop") String prop,
                    @Param("value") String value);

    void setIssueView(TstUser user);
    void setLeftSize(TstUser user);

    List<TstUser> getProjectUsers(@Param("prjId") Integer prjId, @Param("numb") Integer numb);

    void removeFromOrg(@Param("userId") Integer userId,
                       @Param("orgId") Integer orgId);

    void saveIssueColumns(@Param("issueColumns") String issueColumns, @Param("userId") Integer userId);
    void saveIssueFields(@Param("issueFields") String issueFields, @Param("userId") Integer userId);
}
