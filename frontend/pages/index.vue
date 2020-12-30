<template>
  <div>
    <div>TOP</div>
    <div v-if="response.themes.length > 0">
    <table>
      <tr>
        <th>NAME</th>
        <th>DETAIL</th>
      </tr>
      <tr v-for="theme in response.themes" :key="theme.ID">
        <nuxt-link :to="`/themes/${theme.ID}`"><td>{{ theme.Name }}</td></nuxt-link>
        <td>{{ theme.Detail }}</td>
      </tr>
    </table>
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

export default Vue.extend({
  data() {
    return {
      response: {themes: []},
    }
  },
  mounted() {
    axios.get('http://localhost:8080/service/v0/themes')
      .then((response) => {
        this.response = response.data
      })
  }
})
</script>