package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface MsgDao {
    List<TstMsg> query(@Param("userId") Integer userId, @Param("isRead") Boolean isRead);

    void create(TstMsg msg);
}
