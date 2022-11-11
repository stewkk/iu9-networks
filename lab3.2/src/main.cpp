#include <boost/asio.hpp>
#include <boost/asio/spawn.hpp>

#include <thread>

namespace net = boost::asio;
using tcp = net::ip::tcp;
using error_code = boost::system::error_code;

class Session : public std::enable_shared_from_this<Session> {
    public:
        Session(net::io_context& ioc) :
            ioc_(ioc),
            sock_(ioc) {}
        tcp::socket& socket() {
            return sock_;
        }
        void handle() {
            net::spawn(ioc_, std::bind(&Session::do_handle, shared_from_this(), std::placeholders::_1));
        }
    private:
        net::io_context& ioc_;
        tcp::socket sock_;

        void do_handle(net::yield_context yield) {
            char data[128];
            while (true) {
                std::size_t n = sock_.async_read_some(net::buffer(data), yield);
                net::async_write(sock_, net::buffer(data, n), yield);
            }
        }
};


void do_accept(net::yield_context yield, net::io_context& ioc) {
    tcp::endpoint endpoint(tcp::v4(), 2014);
    tcp::acceptor acceptor(ioc, endpoint);
    acceptor.listen();

    while (true) {
        auto session = std::make_shared<Session>(ioc);
        error_code ec;
        acceptor.async_accept(session->socket(), yield[ec]);
        if (!ec) {
            session->handle();
        }
    }
}

std::int32_t main(int argc, char* argv[]) {
    net::io_context ioc;
    net::spawn(ioc, std::bind(do_accept, std::placeholders::_1, std::ref(ioc)));

    ioc.run();
    return 0;
}
