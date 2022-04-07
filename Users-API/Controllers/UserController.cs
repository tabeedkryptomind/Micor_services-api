using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;

namespace Users_API.Controllers;

    [Route("api/[controller]")]
    [ApiController]
    public class UserController : ControllerBase
    {
        private static List<User> usr =  new List<User>
            {
                new User{
                    id = 1,
                    FirstName = "berry",
                    Lastname = "haultman",
                    Email = "hultman@gmail.com",
                    Password = "12412513513"
                }
            };
        [HttpGet]
         public async Task<ActionResult<List<User>>> Get()
         {
            return Ok(usr);
        }
        [HttpPost]
        public async Task<ActionResult<List<User>>> AddUser(User user)
        {
            usr.Add(user);
            return Ok(usr);
        }
    }



