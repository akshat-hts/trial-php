#include <iostream>
using namespace std;

class Transport {
public:
    virtual void deliver() = 0;
};
class Truck : public Transport {
public:
    void deliver() override {
        cout << "By road\n";
    }
};
class Ship : public Transport {
public:
    void deliver() override {
        cout << "By sea\n";
    }
};
class Plane : public Transport {
  public: 
    void deliver() override {
      cout << "By AIR\n"; 
    }
}; 



class Logistics {
public:
    virtual Transport* createTransport() = 0;
    void planDelivery() {
        Transport* t = createTransport();
        t->deliver();
    }
};
class RoadLogistics : public Logistics {
public:
    Transport* createTransport() override {
        return new Truck();
    }
};
class SeaLogistics : public Logistics {
public:
    Transport* createTransport() override {
        return new Ship();
    }
};
class AirLogistics : public Logistics {
  public: 
    Transport * createTransport() override {
      return new Plane(); 
    }
}; 


int main() {
    Logistics* obj1 = new RoadLogistics();
    obj1->planDelivery();
    Logistics * obj2 = new SeaLogistics(); 
    obj2->planDelivery(); 
    Logistics * obj3 = new AirLogistics(); 
    obj3->planDelivery(); 
} 
