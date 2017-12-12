package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestMsgVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {

	@Override
	public List<TestMsg> list(Long userId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestMsg.class);

		if (userId != null) {
			dc.add(Restrictions.eq("runId", userId));
		}

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.asc("pId"));
		dc.addOrder(Order.asc("ordr"));

		List<TestMsg> ls = findAllByCriteria(dc);

		return ls;
	}

    @Override
    public TestMsgVo getById(Long id) {
        TestMsg po = (TestMsg) get(TestMsg.class, id);
        TestMsgVo vo = genVo(po);

        return vo;
    }

    @Override
    public TestMsg save(JSONObject json) {
        String title = json.getString("title");
        String descr = json.getString("descr");
        Long msgId = json.getLong("id");
        Long userId = json.getLong("userId");

        TestMsg msg;
        if (msgId != null) {
            msg = (TestMsg) get(TestMsg.class, msgId);
        } else {
            msg = new TestMsg();
            msg.setUserId(userId);
        }
        msg.setTitle(title);
        msg.setDescr(descr);
        msg.setUserId(userId);
        saveOrUpdate(msg);

        return msg;
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

		return vo;
	}

}

