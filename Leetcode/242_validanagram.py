class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t): 
            return False
        cs, ct = {}, {}
        for i in range(len(s)):
            cs[s[i]] = cs.get(s[i],0) + 1 
            ct[t[i]] = ct.get(t[i],0) + 1 
        for i in cs:
            if cs[i] != ct.get(i,0):
                return False
        return True
        
