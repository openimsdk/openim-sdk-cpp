#define MAX_NUM_OF_CB_S 10
#define MAX_NUM_OF_CB_I_S 10
#define MAX_NUM_OF_CB_S_I_S_S 10
#define MAX_NUM_OF_CB_S_I_S_S_I 10
#define SLEEP_TIME_FOR_GET_INDEX 100 //ms

// use recursive template to generate enough function pointer array
// and define c type function interface
// and define manager class to manage function pool
namespace {
  CB_S* _fps_cb_s=new CB_S[MAX_NUM_OF_CB_S];
  CB_I_S* _fps_cb_i_s=new CB_I_S[MAX_NUM_OF_CB_I_S];
  CB_S_I_S_S* _fps_cb_s_i_s_s=new CB_S_I_S_S[MAX_NUM_OF_CB_S_I_S_S];
  CB_S_I_S_S_I* _fps_cb_s_i_s_s_i=new CB_S_I_S_S_I[MAX_NUM_OF_CB_S_I_S_S_I];
  // c type func interface call cpp function
  std::function<void(const std::string&)>* _cpp_function_cb_s=new std::function<void(const std::string&)>[MAX_NUM_OF_CB_S];
  std::function<void(int,const std::string&)>* _cpp_function_cb_i_s=new std::function<void(int,const std::string&)>[MAX_NUM_OF_CB_I_S];
  std::function<void(const std::string&,int,const std::string&,const std::string&)>* _cpp_function_cb_s_i_s_s=new std::function<void(const std::string&,int,const std::string&,const std::string&)>[MAX_NUM_OF_CB_S_I_S_S];
  std::function<void(const std::string&,int,const std::string&,const std::string&,int)>* _cpp_function_cb_s_i_s_s_i=new std::function<void(const std::string&,int,const std::string&,const std::string&,int)>[MAX_NUM_OF_CB_S_I_S_S_I];

  template<int N>
  void _generate_cb_s(){
    _fps_cb_s[N]=[](char* c_str){
      _cpp_function_cb_s[N](std::string(c_str));
    };
    _generate_cb_s<N-1>();
  }
  template<>
  void _generate_cb_s<0>(){
    _fps_cb_s[0]=[](char* c_str){
      _cpp_function_cb_s[0](std::string(c_str));
    };
  }

  template<int N>
  void _generate_cb_i_s(){
    _fps_cb_i_s[N]=[](int code,char* c_str){
      _cpp_function_cb_i_s[N](code,std::string(c_str));
    };
    _generate_cb_i_s<N-1>();
  }
  template<>
  void _generate_cb_i_s<0>(){
    _fps_cb_i_s[0]=[](int code,char* c_str){
      _cpp_function_cb_i_s[0](code,std::string(c_str));
    };
  }
  template<int N>
  void _generate_cb_s_i_s_s(){
    _fps_cb_s_i_s_s[N]=[](char* operationID,int code,char* c_str,char* c_str2){
      _cpp_function_cb_s_i_s_s[N](std::string(operationID),code,std::string(c_str),std::string(c_str2));
    };
    _generate_cb_s_i_s_s<N-1>();
  }
  template<>
  void _generate_cb_s_i_s_s<0>(){
    _fps_cb_s_i_s_s[0]=[](char* operationID,int code,char* c_str,char* c_str2){
      _cpp_function_cb_s_i_s_s[0](std::string(operationID),code,std::string(c_str),std::string(c_str2));
    };
  }
  template<int N>
  void _generate_cb_s_i_s_s_i(){
    _fps_cb_s_i_s_s_i[N]=[](char* operationID,int code,char* c_str,char* c_str2,int c_int){
      _cpp_function_cb_s_i_s_s_i[N](std::string(operationID),code,std::string(c_str),std::string(c_str2),c_int);
    };
    _generate_cb_s_i_s_s_i<N-1>();
  }
  template<>
  void _generate_cb_s_i_s_s_i<0>(){
    _fps_cb_s_i_s_s_i[0]=[](char* operationID,int code,char* c_str,char* c_str2,int c_int){
      _cpp_function_cb_s_i_s_s_i[0](std::string(operationID),code,std::string(c_str),std::string(c_str2),c_int);
    };
  }

  // init global function pointer array
  void init(){
    _generate_cb_s<MAX_NUM_OF_CB_S-1>();
    _generate_cb_i_s<MAX_NUM_OF_CB_I_S-1>();
    _generate_cb_s_i_s_s<MAX_NUM_OF_CB_S_I_S_S-1>();
    _generate_cb_s_i_s_s_i<MAX_NUM_OF_CB_S_I_S_S_I-1>();
  }
  // define sigle instance class to manage function pool
  class FuncPoolManager{
    private:
    // define a global bitmap, and support atomic operation, to manage cb_s pool
    std::bitset<MAX_NUM_OF_CB_S> _cb_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_I_S> _cb_i_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_S_I_S_S> _cb_s_i_s_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_S_I_S_S_I> _cb_s_i_s_s_i_bitmap;
    std::mutex _cb_s_mutex;
    std::mutex _cb_i_s_mutex;
    std::mutex _cb_s_i_s_s_mutex;
    std::mutex _cb_s_i_s_s_i_mutex;
    FuncPoolManager(){
      init();
    }
    FuncPoolManager(const FuncPoolManager&){}
    public:
    static FuncPoolManager& get_instance(){
      static FuncPoolManager instance;
      return instance;
    }
    // get a available cb_s function index
    int get_cb_s_index(){
      std::lock_guard<std::mutex> lock(_cb_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_bitmap.size();i++){
        if(_cb_s_bitmap[i]==0){
          _cb_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_i_s function index
    int get_cb_i_s_index(){
      std::lock_guard<std::mutex> lock(_cb_i_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_i_s_bitmap.size();i++){
        if(_cb_i_s_bitmap[i]==0){
          _cb_i_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_s_i_s_s function index
    int get_cb_s_i_s_s_index(){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_i_s_s_bitmap.size();i++){
        if(_cb_s_i_s_s_bitmap[i]==0){
          _cb_s_i_s_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_s_i_s_s_i function index
    int get_cb_s_i_s_s_i_index(){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_i_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_i_s_s_i_bitmap.size();i++){
        if(_cb_s_i_s_s_i_bitmap[i]==0){
          _cb_s_i_s_s_i_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // release a available cb_s function index
    int release_cb_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_mutex);
      if(index<0||index>=_cb_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s[index]=nullptr;
      _cb_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_i_s function index
    int release_cb_i_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_i_s_mutex);
      if(index<0||index>=_cb_i_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_i_s[index]=nullptr;
      _cb_i_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_s_i_s_s function index
    int release_cb_s_i_s_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_mutex);
      if(index<0||index>=_cb_s_i_s_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s_i_s_s[index]=nullptr;
      _cb_s_i_s_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_s_i_s_s_i function index
    int release_cb_s_i_s_s_i_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_i_mutex);
      if(index<0||index>=_cb_s_i_s_s_i_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s_i_s_s_i[index]=nullptr;
      _cb_s_i_s_s_i_bitmap[index]=0;
      return 0;
    }
  };
  FuncPoolManager& instance=FuncPoolManager::get_instance();

  // wrapper persistent function
  // wrapp CB_S,if function pool is full,return nullptr
  CB_S _wrapper_cpp_function(const std::function<void(const std::string&)>& cpp_function) {
    int index=FuncPoolManager::get_instance().get_cb_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_index();
    }
    _cpp_function_cb_s[index]=cpp_function;
    return _fps_cb_s[index];
  }
  // wrapp CB_I_S
  CB_I_S _wrapper_cpp_function(const std::function<void(int,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_i_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_i_s_index();
    }
    _cpp_function_cb_i_s[index]=cpp_function;
    return _fps_cb_i_s[index];
  }
  // wrapp CB_S_I_S_S
  CB_S_I_S_S _wrapper_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    }
    _cpp_function_cb_s_i_s_s[index]=cpp_function;
    return _fps_cb_s_i_s_s[index];
  }
  // wrapp CB_S_I_S_S_I
  CB_S_I_S_S_I _wrapper_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&,int)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    }
    _cpp_function_cb_s_i_s_s_i[index]=cpp_function;
    return _fps_cb_s_i_s_s_i[index];
  }

  // wrapp function to onetime function
  CB_S _wrapper_callonce_cpp_function(const std::function<void(const std::string&)>& cpp_function) {
    int index=FuncPoolManager::get_instance().get_cb_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_index();
    }
    _cpp_function_cb_s[index]=[cpp_function,index](const std::string& str)->void {
      cpp_function(str);
      FuncPoolManager::get_instance().release_cb_s_index(index);
    };
    return _fps_cb_s[index];
  }
  
  CB_I_S _wrapper_callonce_cpp_function(const std::function<void(int,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_i_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_i_s_index();
    }
    _cpp_function_cb_i_s[index]=[cpp_function,index](int code,const std::string& str)->void {
      cpp_function(code,str);
      FuncPoolManager::get_instance().release_cb_i_s_index(index);
    };
    return _fps_cb_i_s[index];
  }
  
  CB_S_I_S_S _wrapper_callonce_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    }
    _cpp_function_cb_s_i_s_s[index]=[cpp_function,index](const std::string& operationID,int code,const std::string& str,const std::string& str2)->void {
      cpp_function(operationID,code,str,str2);
      FuncPoolManager::get_instance().release_cb_s_i_s_s_index(index);
    };
    return _fps_cb_s_i_s_s[index];
  }
  
  CB_S_I_S_S_I _wrapper_callonce_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&,int)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    // while loop util get a available index
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    }
    _cpp_function_cb_s_i_s_s_i[index]=[cpp_function,index](const std::string& operationID,int code,const std::string& str,const std::string& str2,int c_int)->void {
      cpp_function(operationID,code,str,str2,c_int);
      FuncPoolManager::get_instance().release_cb_s_i_s_s_i_index(index);
    };
    return _fps_cb_s_i_s_s_i[index];
  }

}

