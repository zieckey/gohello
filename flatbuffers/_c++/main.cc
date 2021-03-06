#include "test_generated.h"
#include <vector>
#include <iostream>
using namespace std;
using namespace TestApp;
int main()
{
    flatbuffers::FlatBufferBuilder builder;
    /////////// Serialize //////////
    // Create list 
    std::vector<uint64_t> vec;
    for(size_t i=0;i<10;i++)
    {
        vec.push_back(i);
    }
    // Create flat buffer inner type
    auto id = 123;
    auto name = builder.CreateString("name");
    auto list = builder.CreateVector(vec); // vector
    auto flag = 1;
    auto kv = KV(1, 1.0); // struct
    // table
    auto mloc = CreateTestObj(builder, id, name, flag, list, &kv);
    builder.Finish(mloc);
    char* ptr = (char*)builder.GetBufferPointer();
    uint64_t size = builder.GetSize();
    
    ////////// Deserialize //////////
    auto obj = GetTestObj((uint8_t*)ptr);
    cout << obj->id() << endl;
    cout << obj->name()->c_str() << endl;
    cout << obj->flag() << endl;
    for(size_t i=0;i<obj->list()->size();i++)
    {
        cout << obj->list()->Get(i) << endl;
    }
    // can use assign to std::vector for speed up
    // vec.reserve(obj->list()->size());
    // vec.assign(obj->list()->begin(), obj->list()->end());
    cout << obj->kv()->key() << endl;
    cout << obj->kv()->value() << endl;
    return 0;
}
