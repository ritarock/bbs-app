<template>
  <div>
    {{ response.data.theme.Name }}
  <br/>
    {{ response.data.theme.Detail }}
  <hr/>
    <div v-for="comment in response.data.comment" :key="comment.ID">
      {{comment.Body}} <br/>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'


type Theme = {
  ID: number
	CreatedAt: string
	UpdatedAt: string
  DeletedAt: string

  Name: string
  Detail: string
  Comments: Comment[]
}
type Comment = {
  ID: number
	CreatedAt: string
	UpdatedAt: string
  DeletedAt: string

  Body: string
  ThemeId: number
}

export default({
    async asyncData({ params }) {
      const response = await axios.get(`http://localhost:8080/service/v0/themes/${params.slug}`)
      return { response }
    }
})

// export default Vue.extend({
//   data() {
//     return {
//       response: {themes: []},
//     }
//   },
//   mounted() {
//     axios.get(`http://localhost:8080/service/v0/themes/${this.$route.params.slug}`)
//       .then((response) => {
//         console.log(response)
//         this.response = response.data
//       })
//   }
// })
</script>