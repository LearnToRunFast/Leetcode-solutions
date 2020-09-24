class Solution:
    def computeArea(self, A: int, B: int, C: int, D: int, E: int, F: int, G: int, H: int) -> int:
        a = (C - A) * (D - B)
        b = (G - E) * (H - F)
        x = max(0, min(C, G) - max( A, E))
        y = max(0, min(D, H) - max(B, F))
        area = a + b - x * y
        return area
