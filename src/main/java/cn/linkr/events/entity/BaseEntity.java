package cn.linkr.events.entity;

import java.io.Serializable;
import java.util.Date;

import javax.persistence.Column;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.MappedSuperclass;
import javax.persistence.Version;


/**
 * <简述>entity基类
 * <详细描述>
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
@MappedSuperclass
public class BaseEntity implements Serializable {

    /**
     *
     */
    private static final long serialVersionUID = -7171038274791404472L;

    /**
     * id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    protected Long id;

    /**
     * 创建时间
     */
    @Column(insertable = true, updatable = false)
    protected Date createTime = new Date();

    /**
     * 更新时间
     */
    @Column(insertable = false, updatable = true)
    protected Date updateTime = new Date();

    /**
     * 是否被删除，状态值
     */
    protected Boolean deleted = Boolean.FALSE;

    /**
     * 是否被禁用，状态值
     */
    protected Boolean disabled = Boolean.FALSE;

    /**
     * 版本控制
     */
    @Version
    protected Integer version;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Date getCreateTime() {
        return createTime;
    }

    public void setCreateTime(Date createTime) {
        this.createTime = createTime;
    }

    public Date getUpdateTime() {
        return updateTime;
    }

    public void setUpdateTime(Date updateTime) {
        this.updateTime = updateTime;
    }

    public Integer getVersion() {
        return version;
    }

    public void setVersion(Integer version) {
        this.version = version;
    }

    public Boolean getDeleted() {
        return deleted;
    }

    public void setDeleted(Boolean deleted) {
        this.deleted = deleted;
    }

    public Boolean getDisabled() {
        return disabled;
    }

    public void setDisabled(Boolean disabled) {
        this.disabled = disabled;
    }
}
