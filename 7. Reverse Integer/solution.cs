public class Solution {
    public int Reverse(int x)
        {
            
            StringBuilder result = new StringBuilder();
            var a = x.ToString().ToArray();
            var isPos = true;
            for (int i = a.Length - 1; i > -1; i--)
            {
                if (a[i] == '-')
                {
                    isPos = false;
                    continue;
                }
                result.Append(a[i]);
            }
            var c = 0;
            int.TryParse(result.ToString(), out c);

            if(!isPos && c!=0)
            {
                return -c;
            }
            return c;
        }
}