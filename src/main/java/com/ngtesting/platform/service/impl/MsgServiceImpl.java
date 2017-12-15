package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestMsgVo;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {

    @Override
    public List<TestMsgVo> list(Long userId, Boolean isRead) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestMsg.class);

        dc.add(Restrictions.eq("userId", userId));
        if (isRead != null) {
            dc.add(Restrictions.eq("isRead", isRead));
        }
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.desc("createTime"));

        List<TestMsg> pos = findAllByCriteria(dc);
        List<TestMsgVo> vos = genVos(pos);

        return vos;
    }

	@Override
	public Page listByPage(Long userId, String isRead,
                           String keywords, Integer currentPage, Integer itemsPerPage) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestMsg.class);

        dc.add(Restrictions.eq("userId", userId));
        if (StringUtil.isNotEmpty(keywords)) {
            dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtils.isNotEmpty(isRead)) {
            dc.add(Restrictions.eq("isRead", Boolean.valueOf(isRead)));
        }
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.desc("createTime"));

        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);

		return page;
	}

    @Override
    public TestMsgVo getById(Long id) {
        TestMsg po = (TestMsg) get(TestMsg.class, id);
        TestMsgVo vo = genVo(po);

        return vo;
    }

    @Override
    public void delete(Long msgId, Long userId) {
        TestMsg po = (TestMsg) get(TestMsg.class, msgId);
        po.setDeleted(true);
        saveOrUpdate(po);
    }

    @Override
    public TestMsg create(TestRun run, Constant.MsgType action, UserVo optUser) {
        TestMsg msg = new TestMsg();

        msg.setName("用户" + StringUtil.highlightDict(optUser.getName()) + action.msg
                + "测试集" + StringUtil.highlightDict(run.getName()));

        msg.setDescr(run.getDescr());
        msg.setUserId(run.getUserId());
        msg.setOptUserId(optUser.getId());
        saveOrUpdate(msg);

        return msg;
    }

    @Override
    public TestMsg markRead(Long id, Long id1) {
        TestMsg po = (TestMsg) get(TestMsg.class, id);
        po.setRead(Boolean.TRUE);
        saveOrUpdate(po);
        return po;
    }

    @Override
    public void markAllReadPers(Long userId) {
        String hql = "update TestMsg msg set msg.isRead=true where msg.userId=? " +
                "AND msg.isRead != true AND msg.deleted != true AND msg.disabled != true";
        getDao().executeByHql(hql, userId);
    }

    @Override
	public List<TestMsgVo> genVos(List<TestMsg> pos) {
        List<TestMsgVo> vos = new LinkedList<>();

        for (TestMsg po: pos) {
			TestMsgVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestMsgVo genVo(TestMsg po) {
		TestMsgVo vo = new TestMsgVo();
		BeanUtilEx.copyProperties(vo, po);

        TestUser user = (TestUser)get(TestUser.class, po.getUserId());
        vo.setAvatar(user.getAvatar());

		return vo;
	}

}

