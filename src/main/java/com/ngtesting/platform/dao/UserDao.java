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

    TstUser get(Integer userId);
    TstUser getByEmail(String nickname);
    TstUser getByEmailAndPassword(@Param("email") String email,
                                  @Param("password") String password);
    TstUser getByToken(@Param("token") String token);
    TstUser getByPhone(@Param("phone") String phone);

    void save(TstUser record);
    void update(TstUser record);

    void setDefaultOrg(@Param("id") Integer id,
                       @Param("orgId") Integer orgId,
                       @Param("orgName") String orgName);

    void setDefaultPrj(@Param("id") Integer id,
                       @Param("prjId") Integer prjId,
                       @Param("prjName") String prjName);

    void setLeftSize(TstUser user);

    List<TstUser> getProjectUsers(@Param("prjId") Integer prjId);

}
