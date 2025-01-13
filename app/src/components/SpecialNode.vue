<script setup lang="ts">

import { computed, onMounted, ref } from 'vue'
import { Position, Handle } from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'
  
const props = defineProps<NodeProps>()

const x = computed(() => `${Math.round(props.position.x)}px`)
const y = computed(() => `${Math.round(props.position.y)}px`)

// reactive state
const count = ref(0)

var status = ref("success")

// functions that mutate state and trigger updates
function increment() {
  count.value++
}

const fetchPosts = async () => {
  try {
    console.log('Fetching new data...');
    const response = await fetch("http://localhost:3000/status");
    const data = await response.json();
    status.value = data.status;
    console.log('Data fetched:', data);
    count.value++;
  } catch(err) {
    console.log('Request failed: ', err.message);
  }
}

// lifecycle hooks
onMounted(() => {
    setInterval(fetchPosts, 5000);

  console.log(`The initial count is ${count.value}.`)
})

</script>

<template>
  <div :class="{ 'vue-flow__node-success': status === 'success', 'vue-flow__node-failure': status === 'failure' }">
    <div>{{ data.label }}</div>

    <div>
        {{ count }}
      {{ x }} {{ y }}
    </div>

    <Handle type="source" :position="Position.Top" />
  </div>
</template>