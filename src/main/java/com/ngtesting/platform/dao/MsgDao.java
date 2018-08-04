package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface MsgDao {
    List<TstMsg> query(@Param("userId") Integer userId,
                       @Param("isRead") Boolean isRead,
                       @Param("keywords") String keywords);

    void create(TstMsg msg);
    TstMsg get(@Param("id") Integer id);

    void markRead(@Param("id") Integer id);

    void markAllRead(@Param("userId") Integer userId);
}
