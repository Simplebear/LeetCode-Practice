public class Solution {
    public int NumUniqueEmails(string[] emails) {
         var resultlist = new List<string>();
            foreach (var email in emails)
            {
                var list = email.Split("@");
                var domain = "@"+list[1];
                var name = list[0].Split("+")[0].Replace(".", "");
                resultlist.Add(name+domain);
            }

            return resultlist.Distinct().Count();
    }
}