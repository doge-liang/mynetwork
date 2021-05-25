package com.graduationProject.dto;

import lombok.Data;

import java.util.Collection;

/**
 * Page
 * <p>
 * 页面类
 * <p>
 * Created : 2021/5/17 15:45
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@Data
public class Page<T extends Collection<?>> {

    private T data;
    private Integer pageSize;
    private Integer totalPage;
    private String bookmark;

    public Page(T data, String bookmark, Integer pageSize){
        this.data = data;
        this.totalPage = data.size();
        this.bookmark = bookmark;
        this.pageSize = 40;
        this.pageSize = pageSize;
    }
}
