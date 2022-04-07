namespace Users_API
{
    public class User 
    {
         public int id { get; set;}
         public string FirstName { get; set; } =  string.Empty;
         public string Lastname {get; set;} = string.Empty;
         public string Email {get; set;} = string.Empty;
         public string Password {get; set;} = string.Empty;  
    }
}