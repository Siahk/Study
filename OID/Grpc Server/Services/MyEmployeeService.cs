using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GrpcServer.Web.Services
{
    public class MyEmployeeService: EmployeeService.EmployeeServiceBase
    {
        private readonly ILogger<MyEmployeeService> _logger;

        public MyEmployeeService(ILooger<MyEmployeeService> logger) 
        {
            _logger = logger;
        }

        public override Task<EmployeeResponse>
            GetByNo(GetByNoRequest request, ServerCallContext context)
            {
                request.No
            }
    }
}