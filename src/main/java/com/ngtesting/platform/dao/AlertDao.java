package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstAlert;
import org.apache.ibatis.annotations.Param;

import java.util.Date;
import java.util.List;

public interface AlertDao {
    List<TstAlert> query(@Param("userId") Integer userId, @Param("isRead") Boolean isRead);

    List<TstAlert> scanAlerts(@Param("userId") Integer userId,
                              @Param("startTimeOfToday") Date startTimeOfToday, @Param("endTimeOfToday") Date endTimeOfToday);

    void create(TstAlert po);

    void markAllRead(@Param("ids") String ids);
}
