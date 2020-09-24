class Codec:

    def serialize(self, root):
        self.vals = []
        def encode(node):
            if node is not None:
                self.vals.append(str(node.val))
                encode(node.left)
                encode(node.right)
            else:
                self.vals.append("#")
        encode(root)
        return ' '.join(self.vals)
    
    def serialize_bfs(self, root):
        self.vals = []
        def encode(root):
            q = [root]
            while q:
                node = q.pop(0)
                if node is not None:
                    self.vals.append(str(node.val))
                    q.append(node.left)
                    q.append(node.right)
                else:
                    self.vals.append('#')
        encode(root)
        return ' '.join(self.vals)
        
    def deserialize(self, data):
        def decode(data):
            val = next(data)
            if val == '#':
                return None
            root = TreeNode(val)
            root.left = decode(data)
            root.right = decode(data)
            return root

        data = iter(data.split())
        return decode(data)
    
    def deserialize_dfs(self, data):
        def decode(vals):
            root = TreeNode(next(vals))
            nodes = [root]
            while nodes:
                node = nodes.pop(0)
                val = next(vals)
                if val != '#':
                    node.left = TreeNode(val)
                    nodes.append(node.left)
                val = next(vals)
                if val != '#':
                    node.right = TreeNode(val)
                    nodes.append(node.right)
            return root

        data = iter(data.split())
        return decode(data)