const TableInfiniteScroll = {
    mounted(el, binding) {
        // 获取element - ui定义好的scroll盒子;
        const SELECTWRAP_DOM = el.querySelector(".el-table__body-wrapper");

        SELECTWRAP_DOM.addEventListener("scroll", function () {
            let sign = 116; // 定义默认的向上滚于乡下滚的边界
            const CONDITION =
                this.scrollHeight - Math.ceil(this.scrollTop) === this.clientHeight; // 注意: && this.scrollTop
            // console.log(this.scrollHeight); // 元素高度
            // console.log(this.scrollTop); // 滚动条高度
            // console.log(this.clientHeight); // 可视部分高度
            // console.log(CONDITION);
            // if (this.scrollTop > sign) {
            //   sign = this.scrollTop;
            //   console.log("向下");
            // }
            // if (this.scrollTop < sign) {
            //   sign = this.scrollTop;
            //   console.log("向上");
            // }

            if (CONDITION) {
                // console.log(binding);
                binding.value();
            }
        });
    }
}

export default TableInfiniteScroll