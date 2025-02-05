Flow State
---

Making a new node:
1. Navigate to functions.ts
2. Add the node to the nodeTypes array and the nodeTypesMap
3. Add the node to the globalCustomNodeTypes
4. Add the additional fields needed to the ConvertNodetoGoNode function
5. Add the nodes additional fields to the convertGoNodeToNode function
6. If applicable, update the updateAllNodesStatus function to update the node status during runtime
7. Create a new node component in the components/nodes folder for your node
8. Navigate to App.vue and update the onNodeDragStop function to add the new node to the nodes array, update it for your new data
9. update the localAddNode function to add the new node to the nodes array
10. Add the template for the new node in the App.vue file (REMEMBER TO IMPORT THE NEW NODE COMPONENT)
11. Build your Node Drawer component in the components/drawers folder
12. Go to the Nodes.go file in the go app and add the new content for the node to the Node struct
13. Add the new node data to the convertRedisDataToNode function
14. Add the new data to the LoadNodes function