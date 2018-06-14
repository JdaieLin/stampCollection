<template>
  <div class="mid-center">
    <div class="user-info">
      <div class="user-head">
        <img src="/static/img/2.png" alt="">
      </div>
      <div class="user-name">{{$store.state.User.name}}</div>
      <div class="user-phone">{{$store.state.User.phone}}</div>
      <div class="user-mail">{{$store.state.User.mail}}</div>
    </div>
    <div class="login-out-btn">
      <div class="action-btn theme-btn white">退出登录</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'User',
  data () {
    return {
    }
  },
  methods: {
  }
}
</script>

<style scoped lang="less">
  @import "../less/common.less";
  .mid-center{
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: #fff;
    .user-info{
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      margin-top: 80px;
      width: 100%;
    }
    .user-head{
      position: relative;
      width: 80px;
      height: 80px;
      border-radius: 50%;
      overflow: hidden;
      margin-bottom: 30px;
      img{
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
    .user-name{
      font-weight: bold;
    }
    .user-phone{
      font-size: 12px;
    }
    .user-mail{
      font-size: 12px;
    }
    .login-out-btn{
      position: absolute;
      width: 100%;
      bottom: 100px;
      padding: 24px 40px;
      .action-btn{
        height: 40px;
        border-radius: 20px;
      }
    }
  }
</style>
