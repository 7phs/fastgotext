#include <iostream>
#include <sstream>
#include <stdio.h>
#include "../fasttext/src/fasttext.h"

const int RES_OK = 0;
const int RES_ERROR_NOT_OPEN = 1;
const int RES_ERROR_WRONG_MODEL = 2;
const int RES_ERROR_NOT_INIT = 3;

extern "C" {
    struct WrapperDictionary {
        std::shared_ptr<const fasttext::Dictionary> dict;
    };

    struct WrapperFastText {
        fasttext::FastText *model;
    };

    struct WrapperVector {
        fasttext::Vector *vector;
    };

    struct PredictRecord {
        float       predict;
        const char* word;
    };

    struct PredictResult {
        PredictRecord* records;
        const char*    err;

        std::vector<PredictRecord> records_;
        std::vector<std::string>   words_;
        std::string    err_;
    };
}

bool checkModelInitialization(const struct WrapperFastText* wrapper) {
    if (wrapper==nullptr ||
        wrapper->model==nullptr ||
        wrapper->model->getDictionary()==nullptr
    ) {
        return false;
    }

    return true;
}

bool checkModelFile(std::istream& in) {
    int32_t magic, version;

    in.read((char*)&(magic), sizeof(int32_t));
    if (magic != FASTTEXT_FILEFORMAT_MAGIC_INT32) {
        return false;
    }

    in.read((char*)&(version), sizeof(int32_t));
    if (version > FASTTEXT_VERSION) {
        return false;
    }

    return true;
}

const int checkVectorsFile(const std::string& path, const int ndim) {
    std::ifstream in(path);

    int64_t n, dim;
    if (!in.is_open()) {
        return RES_ERROR_NOT_OPEN;
    }

    in >> n >> dim;
    if (dim != ndim) {
        return RES_ERROR_WRONG_MODEL;
    }

    in.close();

    return RES_OK;
}

void predictResultResize(struct PredictResult* result, size_t sz) {
    result->records_.resize(sz);
    result->words_.resize(sz);
}

void predictResultSet(struct PredictResult* result, size_t i, std::pair<float, std::string>& rec) {
    auto& new_rec = result->records_[i];
    auto& new_word = result->words_[i];

    new_rec.predict = std::get<0>(rec);
    new_word = std::get<1>(rec);
}

void predictResultSetError(struct PredictResult* result, const char* str) {
    result->err_ = std::string(str);
    result->err = result->err_.c_str();
}

void predictResultFinish(struct PredictResult* result) {
    auto& records = result->records_;
    auto& words = result->words_;

    for(size_t i = 0, sz = records.size(); i < sz; i++) {
        records[i].word = words[i].c_str();
    }

    result->records = records.data();
}

struct WrapperVector* Vector(int ndim) {
    WrapperVector *wrapper = (WrapperVector *)malloc(sizeof (struct WrapperVector));

    wrapper->vector = new fasttext::Vector(ndim);

    return wrapper;
}

struct WrapperDictionary* Dictionary(std::shared_ptr<const fasttext::Dictionary> dict) {
    WrapperDictionary *wrapper = (WrapperDictionary *)malloc(sizeof (struct WrapperDictionary));

    wrapper->dict = dict;

    return wrapper;
}

extern "C" {
    const int DICT_Find(struct WrapperDictionary* wrapper, const char* word) {
        return int(wrapper->dict->getId(word));
    }

    const char* DICT_GetWord(struct WrapperDictionary* wrapper, int id) {
        return wrapper->dict->getWord(id).c_str();
    }

    const int DICT_WordsCount(struct WrapperDictionary* wrapper) {
        return wrapper->dict->nwords();
    }

    void VEC_Release(struct WrapperVector* wrapper) {
        delete wrapper->vector;

        free(wrapper);
    }

    const int VEC_Size(struct WrapperVector* wrapper) {
        return wrapper->vector->size();
    }

    const float* VEC_GetData(struct WrapperVector* wrapper) {
        return wrapper->vector->data_;
    }

    const int PRDCT_Len(struct PredictResult* result) {
        return result->records_.size();
    }

    struct PredictRecord* PRDCT_Records(struct PredictResult* result) {
        return result->records;
    }

    const char* PRDCT_Error(struct PredictResult* result) {
        return result->err;
    }

    void PRDCT_Release(struct PredictResult* result) {
        delete result;
    }

    struct WrapperFastText* FastText() {
        WrapperFastText *wrapper = (WrapperFastText *)malloc(sizeof (struct WrapperFastText));

        wrapper->model = new fasttext::FastText();

        return wrapper;
    }

    const int FT_LoadModel(struct WrapperFastText* wrapper, const char* path) {
        std::ifstream ifs(path, std::ifstream::binary);

        if (!ifs.is_open()) {
            return RES_ERROR_NOT_OPEN;
        }

        if (!checkModelFile(ifs)) {
            return RES_ERROR_WRONG_MODEL;
        }

        wrapper->model->loadModel(ifs);

        ifs.close();

        return RES_OK;
    }

    const int FT_LoadVectors(struct WrapperFastText* wrapper, const char* path) {
        std::string vectorsPath(path);

        if (!checkModelInitialization(wrapper)) {
            return RES_ERROR_NOT_INIT;
        }

        const int res = checkVectorsFile(vectorsPath, wrapper->model->getDimension());
        if (res!=RES_OK) {
            return res;
        }

        wrapper->model->loadVectors(vectorsPath);

        return RES_OK;
    }

    struct WrapperDictionary* FT_GetDictionary(struct WrapperFastText* wrapper) {
        return Dictionary(wrapper->model->getDictionary());
    }

    struct WrapperVector* FT_GetVector(struct WrapperFastText* wrapper, const char* word) {
        struct WrapperVector* wrap_vector = Vector(wrapper->model->getDimension());

        wrapper->model->getWordVector(*wrap_vector->vector, word);

        return wrap_vector;
    }

    struct PredictResult* FT_Predict(struct WrapperFastText* wrapper, const char* text, int k) {
        std::istringstream str(text);
        std::vector<std::pair<float, std::string>> prediction;

        struct PredictResult* result = new struct PredictResult();

        try {
            wrapper->model->predict(str, k, prediction);

            predictResultResize(result, prediction.size());
            for(size_t i = 0, sz = prediction.size(); i<sz; i++) {
                predictResultSet(result, i, prediction[i]);
            }

            predictResultFinish(result);
        } catch(std::exception &e) {
            predictResultSetError(result, e.what());
        }

        return result;
    }

    void FT_Release(struct WrapperFastText* wrapper) {
        delete wrapper->model;

        free(wrapper);
    }
}