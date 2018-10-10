package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstAlert;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IsuFilterDao {
    List<TstAlert> query(@Param("userId") Integer userId, @Param("isRead") Boolean isRead);
}
