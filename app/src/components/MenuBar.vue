<script setup lang="ts">
import { ref, shallowRef, defineEmits, inject, provide } from 'vue'
import PlusIcon from '../assets/Plus.svg'
import ResetIcon from '../assets/Reset.svg'
import AddNodeDrawer from './AddNodeDrawer.vue'

const emit = defineEmits(['addNode', 'reset', 'close-drawer'])
const selectedNode = ref(null)
selectedNode.value = inject('selectedNode')
const showDrawer = ref(false)
const activeIcon = ref(null)

const icons = shallowRef([
  { id: 1, name: 'Add Node', icon: PlusIcon, onClick: () => handleClick(1) },
  { id: 2, name: 'Reset', icon: ResetIcon, onClick: () => handleClick(2) }
])

const hoveredIcon = ref<number | null>(null)

const handleClick = async (id: number) => {
  activeIcon.value = id
  if (id === 1) {

    // send a get request to the server to get the flow state
    const query = {
		Query:      "SELECT * FROM 'test_torq3'",
		Limit:      2,
		Start:      1738040675782000, // 10000 seconds ago
		End:        1738041575782000,
		SourceType: "flowstate",
	}
    const response = await  fetch('http://localhost:3000/query', {
        method: 'POST',
        body: JSON.stringify(query)
    })
    const data = await response.json()
    console.log('Query response:', data)
    showDrawer.value = true
  } else if (id === 2) {
    emit('reset')
  }
}

const closeDrawer = () => {
  showDrawer.value = false
  activeIcon.value = null
  emit('close-drawer')
}

const handleAddNode = (nodeData: { name: string }) => {
  console.log('MenuBar - Received nodeData:', nodeData)
  emit('addNode', nodeData)
  closeDrawer()
}

</script>

<template>
  <div class="dock-container">
    <div class="dock">
      <div
        v-for="icon in icons"
        :key="icon.id"
        class="dock-icon"
        :class="{ 'icon-hovered': hoveredIcon === icon.id }"
        @mouseover="hoveredIcon = icon.id"
        @mouseleave="hoveredIcon = null"
        @click="icon.onClick"
      >
        <div class="icon">
          <component v-if="typeof icon.icon === 'object'" :is="icon.icon" />
          <span v-else>{{ icon.icon }}</span>
        </div>
        <span class="tooltip">{{ icon.name }}</span>
      </div>
    </div>
  </div>

  <!-- Drawer -->
  <AddNodeDrawer 
    v-if="showDrawer || selectedNode.value" 
    @close="closeDrawer"
    @addNode="handleAddNode" 
    :selectedNode="selectedNode.value"
  />
</template>

<style scoped>
.dock-container {
  position: fixed;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  padding-left: 10px;
  z-index: 999;
  pointer-events: auto;
}

.dock {
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  padding: 6px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  transition: padding 0.3s ease;
}

.dock:hover {
  padding: 12px;
}

.dock-icon {
  position: relative;
  transition: all 0.2s ease;
}

.dock:hover .icon {
  margin: 2px 0;
}

.icon {
  width: 35px;
  height: 35px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5em;
  color: white;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.icon-hovered .icon {
  width: 45px;
  height: 45px;
  font-size: 1.8em;
}

.icon-hovered {
  transform: scale(1.2) translateX(10px);
}

.icon-hovered ~ .dock-icon {
  transform: translateY(10px);
}

.dock-icon:hover ~ .dock-icon {
  transform: translateX(10px);
}

.tooltip {
  background: rgba(0, 0, 0, 0.9);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 11px;
  position: absolute;
  left: 100%;
  top: 50%;
  transform: translateY(-50%);
  margin-left: 10px;
  opacity: 0;
  transition: opacity 0.2s ease;
  pointer-events: none;
  white-space: nowrap;
}

.icon-hovered .tooltip {
  opacity: 1;
}

.icon svg {
  width: 20px;
  height: 20px;
  color: white;
  transition: all 0.3s ease;
}

.dock-icon:hover .icon svg {
  width: 26px;
  height: 26px;
}

.icon-hovered .icon svg {
  width: 26px;
  height: 26px;
}

</style>