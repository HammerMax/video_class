<template>
  <div class="app-container">
    <span>vid: {{ videoInfo.vid }} | {{ videoInfo.title }}</span>
    <el-divider></el-divider>
    <el-tabs tab-position="right" style="height: 400px;">
      <el-tab-pane label="详情">
        <el-row>
          <el-col :span="13">
            <video-player
              class="video-player-box"
              ref="videoPlayer"
              :options="playerOptions"
              :playsinline="true"></video-player>
          </el-col>
          <el-col :span="11">
            <el-image style="width:350px; height: 350px" :src="videoInfo.image"></el-image>
            <el-upload
              class="upload-demo"
              action="http://localhost:9999/storage/upload"
              :show-file-list="false"
              :on-success="uploadImgSuccess">
              <el-button size="small" type="primary">点击上传</el-button>
            </el-upload>
          </el-col>
        </el-row>
      </el-tab-pane>
      <el-tab-pane label="截图">截图</el-tab-pane>
    </el-tabs>

    <el-form style="margin-top: 20px" :inline="true" :model="videoInfo" class="demo-form-inline">
      <el-form-item label="标题">
        <el-input v-model="videoInfo.title" ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import 'video.js/dist/video-js.css'
import { videoPlayer } from 'vue-video-player'
import { videoDetail, videoUpdate } from '@/api/video'

export default {
  data() {
    return {
      playerOptions: {
        // videojs options
        height: '400',
        muted: true,
        language: 'en',
        playbackRates: [0.7, 1.0, 1.5, 2.0],
        sources: [{
          type: 'video/mp4',
          src: ''
        }],
        poster: '@/assets/404_images/404.png'
      },
      videoInfo: {
        vid: '',
        image: '',
        title: ''
      },
      staticUrl: 'http://localhost:9999/static'
    }
  },
  components: {
    videoPlayer
  },
  mounted() {
    this.updateVideo()
  },
  methods: {
    uploadImgSuccess(response) {
      if (response.code !== 0) {
        this.$message.error(response.message)
        return
      }
      // 修改图片
      videoUpdate(this.videoInfo.vid, { image: this.staticUrl + '/' + response.data.key }).then(response => {
        console.log(response)
        this.updateVideo()
      })
    },
    updateVideo() {
      videoDetail(this.$route.query.vid).then(response => {
        this.videoInfo = response.data
        this.playerOptions.sources[0].src = this.staticUrl + '/' + this.videoInfo.vid + '.mp4'
      })
    }
  }
}

</script>
