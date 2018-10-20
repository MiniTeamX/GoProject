#include <stdint.h>
#include <iostream>
#include <ostream>
#include <map>
#include <memory> 

template <typename F>
struct ScopeExit {
    ScopeExit(F f) : f(f) {}
    ~ScopeExit() { f(); }
    F f;
};

template <typename F>
ScopeExit<F> MakeScopeExit(F f) {
    return ScopeExit<F>(f);
};

class Header {
public:
    Header(){std::cout << "Header()" <<std::endl;}
    ~Header(){std::cout << "~Header()" <<std::endl;}
};
typedef std::shared_ptr<Header> header_t;

class Route
{
public:
    Route(const header_t &header) : header(header){std::cout << "Route(header_t&)" << std::endl;}
    ~Route(){std::cout << "~Route()" << std::endl;}
    const header_t  header;
};

class RoutingTable
{
private:
    std::map<std::string, Route>    table_;         // peerid -> Route

public:
    void Add(const header_t &header);
    Route* Get();
};

void RoutingTable::Add(const header_t &header)
{
    auto r = Route(header);
    table_.insert(std::pair<std::string, Route>("zhuwen", r));
}

Route* RoutingTable::Get()
{
    std::cout << "enter function RoutingTable::Get()" << std::endl;
    auto it = table_.find("zhuwen");
    if (it == table_.end()) {
        return nullptr;
    }
    return &it->second;
}


RoutingTable table;

header_t GetRouteFromRoutingTable()
{
    header_t header = nullptr;

    auto defer = MakeScopeExit([&header](){
        std::cout << "defer header address:" << &header << std::endl;
        if (header != nullptr) {
            std::cout << "defer found route from routing table" << std::endl;
        } else {
            std::cout << "defer not found route from routing table" << std::endl;
        }    
    });

    table.Add(header_t(new Header()));
    Route *r = table.Get();
    if (r == nullptr) {
        std::cout << "not found route from routing table" << std::endl;
        return nullptr;
    }
    header = r->header;


    std::cout << "header address:" << &header << std::endl;

    if (header != nullptr) {
        std::cout << "found route from routing table" << std::endl;
    } else {
        std::cout << "not found route from routing table" << std::endl;
    }    

    return header;
}

int main(int argc, char **argv) {
    GetRouteFromRoutingTable();
    return 0;
}
