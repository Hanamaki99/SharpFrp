using System.Text;
using System.Runtime.InteropServices;


namespace SharpFrp
{
    internal class Program
    {
        [DllImport("main.dll", EntryPoint = "start_init")]
        extern static int start_init(byte[] test);
        static void Main(string[] args)
        {
            start_init(Encoding.ASCII.GetBytes(args[0]));
        }
    }
}
