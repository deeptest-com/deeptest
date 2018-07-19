package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.MsgDao;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.vo.Page;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {
    @Autowired
    private MsgDao msgDao;

    @Override
    public List<TstMsg> list(Integer userId, Boolean isRead) {

        List<TstMsg> msgs = msgDao.query(userId, false);

        return msgs;
    }

	@Override
	public Page listByPage(Integer userId, String isRead,
                           String keywords, Integer currentPage, Integer itemsPerPage) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstMsg.class);
//
//        dc.add(Restrictions.eq("userId", userId));
//        if (StringUtil.isNotEmpty(keywords)) {
//            dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//        if (StringUtils.isNotEmpty(isRead)) {
//            dc.add(Restrictions.eq("isRead", Boolean.valueOf(isRead)));
//        }
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		dc.addOrder(Order.desc("createTime"));
//
//        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
//
//		return page;
        return null;
	}

    @Override
    public TstMsg getById(Integer id) {
//        TstMsg po = (TstMsg) get(TstMsg.class, id);
//        TstMsg vo = genVo(po);
//
//        return vo;
        return null;
    }

    @Override
    public void delete(Integer msgId, Integer userId) {
//        TstMsg po = (TstMsg) get(TstMsg.class, msgId);
//        po.setDeleted(true);
//        saveOrUpdate(po);
    }

    @Override
    public TstMsg create(TstTask run, Constant.MsgType action, TstUser optUser) {
        TstMsg msg = new TstMsg();

//        msg.setName("用户" + StringUtil.highlightDict(optUser.getName()) + action.msg
//                + "测试集" + StringUtil.highlightDict(run.getName()));
//
//        msg.setDescr(run.getDescr());
//        msg.setUserId(run.getUserId());
//        msg.setOptUserId(optUser.getId());
//        saveOrUpdate(msg);

        return msg;
    }

    @Override
    public TstMsg markReadPers(Integer id, Integer id1) {
//        TstMsg po = (TstMsg) get(TstMsg.class, id);
//        po.setRead(Boolean.TRUE);
//        saveOrUpdate(po);
//        return po;
        return null;
    }

    @Override
    public void markAllReadPers(Integer userId) {
//        String hql = "update TstMsg msg set msg.isRead=true where msg.userId=? " +
//                "AND msg.isRead != true AND msg.deleted != true AND msg.disabled != true";
//        getDao().executeByHql(hql, userId);
    }

}

